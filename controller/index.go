package controller

import (
	"gochat/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// IndexHandler
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func IndexHandler(c *gin.Context) {
	//解析指定文件生成模板对象
	c.HTML(200, "index.html", "index")
}

// ToRegister 跳转到注册页面register.html
func ToRegister(c *gin.Context) {
	c.HTML(200, "/user/register.shtml", nil)
}

// ToChat 跳转到主页面index.html
func ToChat(c *gin.Context) {
	user := models.UserBasic{}
	userID, _ := strconv.Atoi(c.Query("userId"))
	user.ID = uint(userID)
	user.Identity = c.Query("token")
	c.HTML(200, "/chat/index.shtml", user)
}

// InitWebSocket 建立WebSocket连接
func InitWebSocket(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
