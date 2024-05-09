package models

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

// Message 消息
type Message struct {
	gorm.Model
	UserId     uint   //发送者
	TargetId   uint   //接受者
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

var (
	//广播消息发送管道
	udpSendChan chan []byte = make(chan []byte, 1024)
	//读写锁
	rwLocker sync.RWMutex
)

// 发送逻辑
func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			//fmt.Println("sendProc >>>>", string(data))
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
		//fmt.Println("[ws]recvProc <<<<", string(data))
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
		//fmt.Println("udpRecvProc data:", string(buf[:]))
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
			//fmt.Println("udpSendProc:", string(data))
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
	msg.CreateTime = uint64(time.Now().Unix())
	switch msg.Type {
	case 1: //私信
		fmt.Println("dispatch data:", string(data))
		sendMsg(msg.UserId, msg.TargetId, data)
	case 2: //群发
		sendGroupMsg(msg.UserId, msg.TargetId, data)
	case 3: //心跳
		sendHeartBeatMsg(msg.UserId)
	case 4:
	}
}

// 心跳
func sendHeartBeatMsg(userId uint) {
	node := clientMap[userId]
	currentTime := uint64(time.Now().Unix())
	node.HearBeat(currentTime)
}

// 群发
func sendGroupMsg(userId uint, groupId uint, msg []byte) {
	fmt.Println("sendGroupMsg:", string(msg))
	//获取所有群成员id
	contacts := FindContactByGroupId(groupId)
	//向所有群成员发送该消息
	for _, v := range contacts {
		//不能发送给自己
		if userId == v.OwnerId {
			continue
		}
		fmt.Println("sendGroupMsg to: ", v.OwnerId)
		sendMsg(userId, v.OwnerId, msg)
	}
}

// 私聊
func sendMsg(userId uint, targetId uint, msg []byte) {
	fmt.Println("sendMsg to targetId:", targetId, "msg:", string(msg))
	rwLocker.RLock()
	node, ok := clientMap[targetId]
	rwLocker.RUnlock()
	//首先查询用户是否在线，只有在线才会向其发消息
	if GetOnlineUser(targetId) {
		if ok {
			node.DataQueue <- msg
		}
	}
	//缓存所有消息，下次登录时在获取
	SetMessage(userId, targetId, msg)
}

// RedisMsg 缓存消息推送
func RedisMsg(userIdA uint, userIdB uint) []string {
	res, err := GetRedisMsg(userIdA, userIdB)
	if err != nil {
		fmt.Println("没有离线消息")
		return []string{}
	}
	/*//这里不能用sendMsg方法，不然又会去缓存
	for _, v := range r {
		//哪一方发送给哪一方
		msg := Message{}
		json.Unmarshal([]byte(v), &msg)
		userId := msg.UserId
		targetId := msg.TargetId
		rwLocker.RLock()
		node1, _ := clientMap[targetId]
		rwLocker.RUnlock()
		rwLocker.RLock()
		node2, _ := clientMap[userId]
		rwLocker.RUnlock()
		node1.DataQueue <- []byte(v)
		node2.DataQueue <- []byte(v)
	}*/
	//直接返回给前端res
	//fmt.Println(res)
	return res
}
