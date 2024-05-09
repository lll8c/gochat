package models

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"gopkg.in/fatih/set.v0"
)

type Node struct {
	Conn          *websocket.Conn //连接
	Addr          string          //客户端地址
	FirstTime     uint64          //首次连接时间
	HeartbeatTime uint64          //心跳时间
	LoginTime     uint64          //登录时间
	DataQueue     chan []byte     //消息
	GroupSets     set.Interface   //好友/群
}

const (
	HeartBeatMaxTime = 6000 //最大心跳时间
)

var (
	// 存每个用户id与服务器连接节点的映射
	clientMap map[uint]*Node = make(map[uint]*Node, 0)
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
	userId, _ := strconv.Atoi(query.Get("userId"))
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
	currentTime := uint64(time.Now().Unix())
	node := &Node{
		Conn:          conn,
		Addr:          conn.RemoteAddr().String(),
		FirstTime:     currentTime,
		HeartbeatTime: currentTime,
		LoginTime:     currentTime,
		DataQueue:     make(chan []byte, 50),
		GroupSets:     set.New(set.ThreadSafe),
	}
	rwLocker.Lock()
	clientMap[uint(userId)] = node
	rwLocker.Unlock()
	//用户关系
	//fmt.Println("userId=", userId)
	//发送和接收逻辑
	go sendProc(node)
	go recvProc(node)
	//加入在线用户到缓存
	SetUserOnlineInfo(userId, []byte(node.Addr), time.Duration(viper.GetInt("timeout.RedisOnlineTime"))*time.Hour)
	node.DataQueue <- []byte("欢迎进入聊天系统")
}

// HearBeat 更新用户心跳
func (node *Node) HearBeat(currentTime uint64) {
	node.HeartbeatTime = currentTime
	return
}

// IsHeartBeatTimeOut 判断用户心跳是否超时
func (node *Node) IsHeartBeatTimeOut(currentTime uint64) (timeOut bool) {
	if node.HeartbeatTime+viper.GetUint64("timeout.HeartBeatMaxTime") <= currentTime {
		timeOut = true
	}
	return
}

// CleanConnection 清理超时连接
func CleanConnection(param interface{}) (result bool) {
	result = true
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("cleanConnection err")
		}
	}()
	fmt.Println("定时任务，清理超时连接", param)
	currenTime := uint64(time.Now().Unix())
	for id, node := range clientMap {
		if node.IsHeartBeatTimeOut(currenTime) {
			fmt.Println(id, "心跳超时，自动下线关闭连接")
			node.Conn.Close()
			//将用户从缓存中删除
			DeleteUserOnline(id)
		}
	}
	return result
}
