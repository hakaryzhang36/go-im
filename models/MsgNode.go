package models

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// 通信节点，相当于用户发消息的管道一头的节点
type MsgNode struct {
	Id             string
	Conn           *websocket.Conn // 管道
	WriteDataQueue chan []byte     // 写缓冲区，发送到用户的消息
	ReadDataQueue  chan []byte     // 读缓冲区，从用户发来的消息
	//GroupSets set.Interface
}

func (table *MsgNode) TableName() string {
	return "message_node"
}

// 关闭节点
func (node *MsgNode) CloseClient() {
	defer func() {
		err := recover()
		if nil != err {
			log.Println("关闭节点出错, id=", node.Id)
		}
	}()
	if _, ok := clientMap[node.Id]; ok {
		node.Conn.Close()
		rwLocker.Lock()
		delete(clientMap, node.Id)
		rwLocker.Unlock()
		log.Println("关闭节点, id=", node.Id)
	}
}

// 映射关系，用于记录用户拥有的节点
var clientMap map[string]*MsgNode = make(map[string]*MsgNode, 0)

// 读写锁
var rwLocker sync.RWMutex

func InitNode(node *MsgNode) {
	// 【TODO】将未读缓存加载进 Node
	rwLocker.Lock()
	clientMap[node.Id] = node
	rwLocker.Unlock()

	// 启动协程进行放送、读取
	go sendProc(node)
	go recvProc(node)
	go recvHandlerData(node)
}

// 发送协程，从写缓冲区中获取消息，通过ws发送
func sendProc(node *MsgNode) {
	defer node.CloseClient()
	for data := range node.WriteDataQueue {
		err := node.Conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			fmt.Println("[WS] WriteMessage error: ", err)
			return
		}
	}
}

// 接收协程，从ws连接中接收消息，并写入读缓冲区
func recvProc(node *MsgNode) {
	defer node.CloseClient()
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println("[WS] ReadMessage error: ", err)
			return
		}
		node.ReadDataQueue <- data // 将消息压入读缓冲区，交由其他协程处理
	}
}

// 消息接收后的后端处理逻辑
func recvHandlerData(node *MsgNode) {
	defer node.CloseClient()
	for {
		select {
		case data := <-node.ReadDataQueue:
			fmt.Println("发送消息: ", data)
			go dispatch(data) // ?
		case <-time.NewTimer(100 * time.Second).C:
			// 【TD】超时后处理
			// 把写缓存中的消息全部发送出去
			// 读缓存中消息全部存储
			fmt.Println("节点长时间没响应断开, ", node)
			return
		}

	}
}

func dispatch(data []byte) {
	msg := &Message{}
	msg.Content = string(data)
	msg.RecvId = "222" // 接收方id
	msg.Type = 1
	// err := json.Unmarshal(data, msg)
	// if err != nil {
	// 	fmt.Println("Msg 格式错误: ", err)
	// 	return
	// }
	switch msg.Type {
	case 1:	// 一对一私聊
		sendMsg(msg.RecvId, data)

	}
}

func sendMsg(recvId string, data []byte) {
	rwLocker.RLock()
	node, ok := clientMap[recvId]
	rwLocker.RUnlock()
	if ok {
		node.WriteDataQueue <- data
	} else {
		fmt.Println("连接节点不存在: ", recvId)
	}
}
