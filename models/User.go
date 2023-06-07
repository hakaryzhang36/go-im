package models

import (
	"fmt"

	"gorm.io/gorm"
	"zhangteam.org/im-project/utils"
)

type User struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     uint64
	HeartbeatTime uint64
	LogoutTime    uint64
	IsLogout      bool
	DeviceInfo    string
}

func (table *User) TableName() string {
	return "user"
}

func GetUserList() []*User {
	userList := make([]*User, 10)
	utils.DB.Find(&userList)
	for _, v := range userList {
		fmt.Println(v)
	}
	return userList
}

func CreateUser(user User) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeteleUser(user User) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user User) *gorm.DB {
	return utils.DB.Model(&user).Updates(User{Name: user.Name, Password: user.Password})
}

func FindUserByName(name string) []User {
	userList := []User{}
	utils.DB.Where(&User{Name: name}).Find(&userList)
	return userList
}

func FindUserByEmail(email string) []User {
	userList := []User{}
	utils.DB.Where(&User{Email: email}).Find(&userList)
	return userList
}

func FindUserByPhone(phone string) []User {
	userList := []User{}
	utils.DB.Where(&User{Phone: phone}).Find(&userList)
	return userList
}
