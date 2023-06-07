package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"zhangteam.org/im-project/utils"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// @BasePath
// SendMsg
// @Summary 发送消息
// @Schemes
// @Tags msgService
// @Param msg formData string false "Message"
// @Param sendUser formData string false "Sender Id"
// @Param recvUser formData string false "Reciver Id"
// @Produce json
// @Success 200 {string} success
// @Router /msg/sendMsg [post]
func SendMsg(c *gin.Context) {
	// msg, _ := strconv.Atoi(c.PostForm("msg"))

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		defer conn.Close()
		for {
			_, p, err := conn.ReadMessage()
			fmt.Println("Send msg: ", string(p))
			if err != nil {
				log.Println(err)
				return
			} else {
				utils.Publish(c, "channel-test", string(p))
			}
		}
	}()

	go func() {
		defer conn.Close()
		for {
			bmsg, _ := utils.Subscribe(c, "channel-test")
			back := []byte(bmsg)
			if err := conn.WriteMessage(1, back); err != nil {
				log.Println(err)
			}
		}
	}()

}
