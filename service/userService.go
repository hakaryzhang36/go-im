package service

import (
	"fmt"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"zhangteam.org/im-project/models"
	"zhangteam.org/im-project/utils"
)

// @BasePath

// GetUserList
// @Summary 获取所有用户
// @Schemes
// @Description get all users.
// @Tags userService
// @Produce json
// @Success 200 {string} success
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(200, gin.H{
		"message":  "success",
		"userList": data,
	})
}

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.User{}

	//name := c.Query("name")
	//password := c.Query("password")
	name := c.Request.FormValue("name")
	password := c.Request.FormValue("password")
	fmt.Println(models.FindUserByName(name))
	user := models.FindUserByName(name)[0]
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "该用户不存在",
			"data":    data,
		})
		return
	}

	flag := (password == user.Password) // utils.ValidPassword(password, user.Salt, user.Password)
	if !flag {
		c.JSON(200, gin.H{
			"code":    -1, //  0成功   -1失败
			"message": "密码不正确",
			"data":    data,
		})
		return
	}
	// pwd := utils.MakePassword(password, user.Salt)
	data = user

	c.JSON(200, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "登录成功",
		"data":    data,
	})
}

// @BasePath

// CreateUser
// @Summary 新建用户
// @Schemes
// @Tags userService
// @Param name formData string false "User Name"
// @Param password formData string false "Password"
// @Param password_repeat formData string false "Repeat Password"
// @Param email formData string false "Email"
// @Param phone formData string false "Phone"
// @Produce json
// @Success 200 {string} success
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	user := models.User{}
	user.Name = c.PostForm("name")
	password := c.PostForm("password")
	password_repeat := c.PostForm("password_repeat")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if password != password_repeat {
		c.JSON(400, gin.H{
			"message": "repeat password wrong!",
		})
		return
	}

	if len(models.FindUserByName(user.Name)) != 0 {
		c.JSON(400, gin.H{
			"message": "name has been register!",
		})
		return
	}

	// if len(models.FindUserByEmail(user.Email)) != 0 {
	// 	c.JSON(400, gin.H{
	// 		"message": "email has been register!",
	// 	})
	// 	return
	// }

	// if len(models.FindUserByPhone(user.Phone)) != 0 {
	// 	c.JSON(400, gin.H{
	// 		"message": "phone has been register!",
	// 	})
	// 	return
	// }

	user.Password = password
	_ = models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    "200",
		"message": "success",
	})
}

// @BasePath
// DeleteUser
// @Summary 删除用户
// @Schemes
// @Tags userService
// @Param id formData string false "User Id"
// @Produce json
// @Success 200 {string} success
// @Router /user/deleteUser [post]
func DeleteUser(c *gin.Context) {
	user := models.User{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	_ = models.DeteleUser(user)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

// @BasePath
// UpdateUser
// @Summary 更新用户信息
// @Schemes
// @Tags userService
// @Param id formData string false "User Id"
// @Param name formData string false "Name"
// @Param password formData string false "Password"
// @Produce json
// @Success 200 {string} success
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.User{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	_ = models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func SearchFriends(c *gin.Context) {
	id, _ := strconv.Atoi(c.Request.FormValue("userId"))
	users := models.SearchFriend(uint(id))
	// c.JSON(200, gin.H{
	// 	"code":    0, //  0成功   -1失败
	// 	"message": "查询好友列表成功！",
	// 	"data":    users,
	// })
	utils.RespOKList(c.Writer, users, len(users))
}
