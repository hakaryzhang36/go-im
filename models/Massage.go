package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SendId  string
	RecvId  string
	Type    int
	Media   int
	Content string
	Pic     string
	Url     string
	Desc    string
	Amount  int
}

func (table *Message) TableName() string {
	return "message"
}

// 通信节点，相当于用户发消息的管道一头的节点
type Node struct {
	Id             string
	Conn           *websocket.Conn // 管道
	WriteDataQueue chan []byte     // 写缓冲区
	ReadDataQueue  chan []byte     // 读缓冲区
	//GroupSets set.Interface
}

// 映射关系，用于记录用户拥有的节点
var clientMap map[string]*Node = make(map[string]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

// 用户进入系统后，生成后续连接的资源
func Chat(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id := query.Get("id")
	isValida := true // checkToken()

	conn, err := (&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// token check
		CheckOrigin: func(r *http.Request) bool {
			return isValida
		},
	}).Upgrade(writer, request, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	// 创建通信节点，并绑定到用户
	node := &Node{
		Id:             id,
		Conn:           conn,
		WriteDataQueue: make(chan []byte, 1024),
		ReadDataQueue:  make(chan []byte, 1024),
	}
	// 【TODO】将未读缓存加载进 Node
	rwLocker.Lock()
	clientMap[id] = node
	rwLocker.Unlock()

	// 启动协程进行放送、读取
	go sendProc(node)
	go recvProc(node)
	go recvHandlerData(node)

}

// 发送协程，对通信节点循环操作，将消息写入写缓冲区
func sendProc(node *Node) {
	defer node.CloseClient()
	for data := range node.WriteDataQueue {
		err := node.Conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			fmt.Println("[WS] WriteMessage error: ", err)
		}
	}
}

// 接收协程，对通信节点循环操作，从ws连接中接收消息，并写入读缓冲区
func recvProc(node *Node) {
	defer node.CloseClient()
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println("[ws] ReadMessage error: ", err)
		}
		node.ReadDataQueue <- data // 将消息压入读缓冲区，交由其他协程处理
	}
}

// 消息接收后的后端处理逻辑
// 从读缓冲区中读取消息，并发送给目标ip
func recvHandlerData(node *Node) {
	defer node.CloseClient()
	for {
		select {
		case data := <-node.ReadDataQueue:
			fmt.Println("发送消息: ", data)
			go dispatch(data) // ?
		case <-time.NewTimer(10 * time.Second).C:
			// 【TD】超时后处理
			// 把写缓存中的消息全部发送出去
			// 读缓存中消息全部存储
			fmt.Println("节点长时间没响应断开, ", node)
			return
		}

	}
}

// 关闭连接
func (node *Node) CloseClient() {
	defer func() {
		err := recover()
		if nil != err {
			log.Println("移除用户出错:用户id=", node.Id)
		}
	}()
	if _, ok := clientMap[node.Id]; ok {
		node.Conn.Close()
		rwLocker.Lock()
		delete(clientMap, node.Id)
		rwLocker.Unlock()
		log.Println("移除用户,用户id=", node.Id)
	}
}

func dispatch(data []byte) {
	msg := &Message{}
	err := json.Unmarshal(data, msg)
	if err != nil {
		fmt.Println("Msg 格式错误: ", err)
		return
	}
	switch msg.Type {
	case 1:
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
