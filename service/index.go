package service

import (
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
	"zhangteam.org/im-project/models"
)

// @BasePath

// GetIndex
// @Summary 首页
// @Schemes
// @Description 访问首页的接口
// @Tags index
// @Produce json
// @Success 200 {string} welcome!
// @Router /index [get]
func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles("index.html", "views/chat/head.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "index")
	// c.JSON(200, gin.H{
	// 	"message": "welcome !!  ",
	// })
}

// @BasePath

// ToRegister
// @Summary 注册页面
// @Schemes
// @Description 访问首页的接口
// @Tags index
// @Produce json
// @Success 200 {string} welcome!
// @Router /toRegister [get]
func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "register")
	// c.JSON(200, gin.H{
	// 	"message": "welcome !!  ",
	// })
}

func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles("views/chat/index.html",
		"views/chat/head.html",
		"views/chat/foot.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/createcom.html",
		"views/chat/userinfo.html",
		"views/chat/main.html")
	if err != nil {
		panic(err)
	}
	userId, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user := models.User{}
	user.ID = uint(userId)
	user.Identity = token
	//fmt.Println("ToChat>>>>>>>>", user)
	ind.Execute(c.Writer, user)
	// c.JSON(200, gin.H{
	// 	"message": "welcome !!  ",
	// })
}
