package models

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
)

// Message 消息
type Message struct {
	gorm.Model
	UserId     int64  //发送者
	TargetId   int64  //接受者
	Type       int    //发送类型  1私聊  2群聊  3心跳
	Media      int    //消息类型  1文字 2表情包 3语音 4图片/表情包
	Content    string //消息内容
	CreateTime uint64 //创建时间
	ReadTime   uint64 //读取时间
	Pic        string
	Url        string
	Desc       string
	Amount     int //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

var (
	// 存每个用户id与服务器连接节点的映射
	clientMap map[int64]*Node = make(map[int64]*Node, 0)
	//广播消息发送管道
	udpSendChan chan []byte = make(chan []byte, 1024)
	//读写锁
	rwLocker sync.RWMutex
)

// 广播发送和接收协程
func init() {
	go udpSendProc()
	go udpRecvProc()
}

// Chat 需要：发送者ID ，接受者ID ，消息类型，发送的内容，发送类型
func Chat(writer http.ResponseWriter, request *http.Request) {
	//1.获取参数并检验token等合法性
	query := request.URL.Query()
	userId, _ := strconv.ParseInt(query.Get("userId"), 10, 64)
	isvalida := true //checkToke()
	//升级为websocket连接，进行token校验
	conn, err := (&websocket.Upgrader{
		//token 校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	//2.保存连接
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()
	//用户关系
	//fmt.Println("userId=", userId)
	//发送和接收逻辑
	go sendProc(node)
	go recvProc(node)
	//加入在线用户到缓存
	//SetUserOnlineInfo("online_"+Id, []byte(node.Addr), time.Duration(viper.GetInt("timeout.RedisOnlineTime"))*time.Hour)
	sendMsg(userId, []byte("欢迎进入聊天系统"))
}

// 发送逻辑
func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			fmt.Println("sendProc >>>>", string(data))
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println("writeMessage err:", err)
				return
			}
		}
	}
}

// 接收逻辑
func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println("readMessage err:", err)
			return
		}
		//广播发送消息给所有用户
		udpSendChan <- data
		fmt.Println("[ws]recvProc <<<<", string(data))
	}
}

/*func broadMsg(data []byte) {
	udpSendChan <- data
}*/

// 完成udp数据接收
func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 9999,
	})
	if err != nil {
		fmt.Println("net.listenUDP err:", err)
		return
	}
	defer con.Close()
	//读取数据
	for {
		var buf [512]byte
		n, err := con.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("udpRecvProc data:", string(buf[:]))
		dispatch(buf[:n])
	}
}

// 完成udp数据发送
func udpSendProc() {
	//走路由网关地址
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(192, 168, 83, 255),
		Port: 9999,
	})
	if err != nil {
		fmt.Println("net.dialUDP err:", err)
		return
	}
	defer con.Close()

	//发送数据
	for {
		select {
		case data := <-udpSendChan:
			fmt.Println("udpSendProc:", string(data))
			_, err := con.Write(data)
			if err != nil {
				fmt.Println("con.write err:", err)
				return
			}
		}
	}
}

// 服务器收到消息后，后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1: //私信
		fmt.Println("dispatch data:", string(data))
		sendMsg(msg.TargetId, data)
	case 2: //群发
		//sendGroupMsg()
	case 3:
		//sendAllMsg()
	case 4:
	}
}

// 私聊
func sendMsg(targetId int64, msg []byte) {
	fmt.Println("sendMsg >>> targetId:", targetId, "msg:", string(msg))
	rwLocker.RLock()
	node, ok := clientMap[targetId]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
