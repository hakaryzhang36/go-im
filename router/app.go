package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "zhangteam.org/im-project/docs"
	"zhangteam.org/im-project/service"
)

func Router() *gin.Engine {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.GET("/toChat", service.ToChat)
	r.POST("/searchFriends", service.SearchFriends)

	r.GET("/user/getUserList", service.GetUserList)
	r.POST("/user/createUser", service.CreateUser)
	r.POST("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.GET("/msg/SendUserMsg", service.SendUserMsg)
	r.POST("user/findUserByNameAndPwd", service.FindUserByNameAndPwd)
	r.GET("/toRegister", service.ToRegister)

	r.Static("/asset", "asset/")
	r.LoadHTMLGlob("views/**/*")
	return r
}
