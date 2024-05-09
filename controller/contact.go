package controller

import (
	"gochat/models"
	"gochat/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddFriendById
// @Summary 添加好友
// @Tags 用户模块
// @param userId formData string false "userID"
// @param targetId formData string false "targetID"
// @Success 200 {string} json{"code","message"}
// @Router /contact/addFriend [post]
func AddFriendById(c *gin.Context) {
	//获取userID和targetID
	userID, _ := strconv.Atoi(c.PostForm("userID"))
	targetID, _ := strconv.Atoi(c.PostForm("targetID"))
	if userID == targetID {
		utils.RespFail(c.Writer, "不能添加自己为好友")
		return
	}
	//添加数据库库
	code, msg := models.AddFriend(uint(userID), uint(targetID))
	if code == 0 {
		utils.RespOK(c.Writer, "添加好友成功", code)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}

func AddFriendByName(c *gin.Context) {
	//获取userID和targetID
	userID, _ := strconv.Atoi(c.PostForm("userID"))
	targetName := c.PostForm("targetName")
	if targetName == "" {
		utils.RespFail(c.Writer, "输入不能为空")
		return
	}
	user := models.FindUserByName(targetName)
	if uint(userID) == user.ID {
		utils.RespFail(c.Writer, "不能添加自己为好友")
		return
	}
	//添加数据库库
	code, msg := models.AddFriend(uint(userID), user.ID)
	if code == 0 {
		utils.RespOK(c.Writer, "添加好友成功", code)
	} else {
		utils.RespFail(c.Writer, msg)
	}
}
