package router

import (
	"gochat/controller"
	_ "gochat/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//swagger
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	//加载静态资源
	r.Static("/asset", "asset/")
	//解析文件
	r.LoadHTMLGlob("view/**/*")
	//首页
	r.GET("/", controller.IndexHandler)
	r.GET("/index", controller.IndexHandler)
	r.GET("/toRegister", controller.ToRegister)
	r.GET("/toChat", controller.ToChat)
	r.POST("/searchFriends", controller.SearchFriends)
	r.GET("/initWebSocket", controller.InitWebSocket)

	//用户模块
	r.GET("/user/getUserList", controller.GetUserList)
	r.POST("/user/createUser", controller.CreateUser)
	r.GET("/user/deleteUser", controller.DeleteUser)
	r.POST("/user/updateUser", controller.UpdateUser)
	r.POST("/user/findUserByNameAndPwd", controller.FindUserByNameAndPwd)

	//发送消息
	r.GET("/user/sendMsg", controller.SendMsg)
	r.GET("/user/sendUserMsg", controller.SendUserMsg)
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	return r
}