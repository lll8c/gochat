package controller

import (
	"fmt"
	"gochat/models"
	"gochat/utils"
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

// SearchFriends 进入chat主页时，查询所有好友
func SearchFriends(c *gin.Context) {
	str := c.PostForm("userId")
	userID, _ := strconv.Atoi(str)
	//fmt.Println(userID)
	users := models.SearchFriend(uint(userID))
	utils.RespOKList(c.Writer, users, len(users))
}

// LoadCommunity 进入chat主页时，加载所有群聊
func LoadCommunity(c *gin.Context) {
	str := c.PostForm("userId")
	userID, _ := strconv.Atoi(str)
	data := models.SearchCommunity(uint(userID))
	if len(data) == 0 {
		utils.RespFail(c.Writer, "暂无群聊")
	} else {
		utils.RespOKList(c.Writer, data, len(data))
	}
	fmt.Println(data)
}
