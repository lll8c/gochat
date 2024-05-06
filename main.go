package main

import (
	"gochat/models"
	"gochat/router"
	"gochat/utils"
	"time"

	"github.com/spf13/viper"
)

func main() {
	utils.InitConfig("conf/config.yaml")
	utils.InitMySQL()
	utils.InitRedis()
	InitTimer()
	//自动迁移创建表或添加字段
	//utils.DB.AutoMigrate(&models.UserBasic{})
	//utils.DB.AutoMigrate(&models.Message{})
	//utils.DB.AutoMigrate(&models.Contact{})
	//utils.DB.AutoMigrate(&models.GroupBasic{})
	//utils.DB.AutoMigrate(&models.Community{})
	r := router.Router()
	r.Run()
}

// InitTimer 初始化定时器定时清理超时连接
func InitTimer() {
	utils.Timer(time.Duration(viper.GetInt("timeout.DelayHeartBeatTime"))*time.Second, time.Duration(viper.GetInt("timeout.HeartBeatHZ"))*time.Second, models.CleanConnection, "")
}
