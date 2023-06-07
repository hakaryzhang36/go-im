package main

import (
	"zhangteam.org/im-project/router"
	"zhangteam.org/im-project/test"
	"zhangteam.org/im-project/utils"
)

func main() {
	unitTest()

}

func initServer() {
	utils.InitConfig()
	utils.InitDB()
	utils.InitRedis()

	r := router.Router()
	r.Run("localhost:8080")
}

func unitTest() {
	test.TestGorm()
}
