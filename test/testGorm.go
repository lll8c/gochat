package main

import (
	"fmt"
	"gochat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123@tcp(127.0.0.1:3306)/gochat?charset=utf8mb4&parseTime=true&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("gorm.open err:", err)
		return
	}
	//自动迁移 根据模型自动添加字段
	db.AutoMigrate(&models.UserBasic{})

	//创建
	user := &models.UserBasic{}
	user.Name = "zs"
	db.Create(user)

	//查询
	db.First(user, 1)
	fmt.Println(user) // 根据整型主键查找

	//修改
	db.Model(user).Update("PassWord", "1234")
	//Update - 更新多个字段
	db.Model(user).Updates(models.UserBasic{Name: "ls"}) // 仅更新非零值字段
	db.Model(user).Updates(map[string]interface{}{"name": "ls"})

	//删除
	//db.Delete(user, 1)
}
