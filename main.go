package main

import (
	"gochat/router"
	"gochat/utils"
)

func main() {
	utils.InitConfig("conf/config.yaml")
	utils.InitMySQL()
	utils.InitRedis()
	//自动迁移创建表或添加字段
	//utils.DB.AutoMigrate(&models.UserBasic{})
	//utils.DB.AutoMigrate(&models.Message{})
	//utils.DB.AutoMigrate(&models.Contact{})
	//utils.DB.AutoMigrate(&models.GroupBasic{})
	//utils.DB.AutoMigrate(&models.Community{})
	r := router.Router()
	r.Run()
}
