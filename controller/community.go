package controller

import (
	"gochat/models"
	"gochat/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateCommunity
// @Summary 创建群
// @Tags 用户模块
// @param ownerId formData string false "ownerId"
// @param name formData string false "name"
// @Success 200 {string} json{"code","message"}
// @Router /contact/createCommunity [post]
func CreateCommunity(c *gin.Context) {
	//获取userID和targetID
	ownerId, _ := strconv.Atoi(c.PostForm("ownerId"))
	name := c.PostForm("name")
	desc := c.PostForm("desc")
	community := models.Community{
		OwnerId: uint(ownerId),
		Name:    name,
		Desc:    desc,
	}
	//添加数据库库
	code, msg := models.CreateCommunity(community)
	if code == 0 {
		utils.RespOK(c.Writer, msg, code)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}

// JoinGroup
// @Summary 加入群聊
// @Tags 用户模块
// @param userId formData string false "userId"
// @param comId formData string false "comId"
// @Success 200 {string} json{"code","message"}
// @Router /contact/joinGroup [post]
func JoinGroup(c *gin.Context) {
	//获取userID和targetID
	userId, _ := strconv.Atoi(c.PostForm("userId"))
	groupId, _ := strconv.Atoi(c.PostForm("comId"))
	//添加到数据库
	code, msg := models.JoinGroup(uint(userId), uint(groupId))
	if code == 0 {
		utils.RespOK(c.Writer, msg, code)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}
