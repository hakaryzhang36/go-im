package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Rdb *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}

func InitDB() {
	add := viper.GetString("mysql.address")
	port := viper.GetString("mysql.port")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	database := viper.GetString("mysql.database")
	charset := viper.GetString("mysql.charset")
	parseTime := viper.GetString("mysql.parseTime")
	loc := viper.GetString("mysql.loc")
	dns := username + ":" + password + "@tcp(" + add + ":" + port + ")/" + database + "?charset=" + charset + "&parseTime=" + parseTime + "&loc=" + loc

	// 自定义数据库日志，实现控制台打印log
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	DB, _ = gorm.Open(mysql.Open(dns), &gorm.Config{Logger: newLogger})
}

func InitRedis() {
	ctx := context.Background()
	Rdb = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr") + ":" + viper.GetString("redis.port"),
		PoolSize: viper.GetInt("redis.poolSize"),
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.db"),          // use default DB
	})

	_, err := Rdb.Ping(ctx).Result()
	if err == nil {
		fmt.Println("[Redis] Server Connected, on", Rdb.Options().Addr)
	} else {
		fmt.Println(err)
	}
}
