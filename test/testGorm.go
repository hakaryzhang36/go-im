package test

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zhangteam.org/im-project/models"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestGorm() {

	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/imdb?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema （在数据库中生成对应的表）
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.Group{})

	// Create
	// user := &models.UserBasic{}
	// user.Name = "t1"
	// db.Create(user)

	// // Read
	// var u models.UserBasic
	// db.First(&u, 1) // 根据整型主键查找
	// fmt.Println(u)

	// // Update - 将 product 的 price 更新为 200
	// db.Model(&u).Update("Password", "123456")
}
