package models

import (
	"fmt"

	"gorm.io/gorm"
	"zhangteam.org/im-project/utils"
)

type UserBasic struct {
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

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	userList := make([]*UserBasic, 10)
	utils.DB.Find(&userList)
	for _, v := range userList {
		fmt.Println(v)
	}
	return userList
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeteleUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, Password: user.Password})
}

func FindUserByName(name string) []UserBasic {
	userList := []UserBasic{}
	utils.DB.Where(&UserBasic{Name: name}).Find(&userList)
	return userList
}

func FindUserByEmail(email string) []UserBasic {
	userList := []UserBasic{}
	utils.DB.Where(&UserBasic{Email: email}).Find(&userList)
	return userList
}

func FindUserByPhone(phone string) []UserBasic {
	userList := []UserBasic{}
	utils.DB.Where(&UserBasic{Phone: phone}).Find(&userList)
	return userList
}
