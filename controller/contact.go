package controller

import (
	"gochat/models"
	"gochat/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddFriend
// @Summary 添加好友
// @Tags 用户模块
// @param userId formData string false "userID"
// @param targetId formData string false "targetID"
// @Success 200 {string} json{"code","message"}
// @Router /contact/addFriend [post]
func AddFriend(c *gin.Context) {
	//获取userID和targetID
	userID, _ := strconv.Atoi(c.PostForm("userID"))
	targetID, _ := strconv.Atoi(c.PostForm("targetID"))
	if userID == targetID {
		utils.RespFail(c.Writer, "不能添加自己为好友")
		return
	}
	//添加数据库库
	code := models.AddFriend(uint(userID), uint(targetID))
	if code == 0 {
		utils.RespOK(c.Writer, "添加好友成功", code)
	} else if code == 1 {
		utils.RespFail(c.Writer, "已添加该用户为好友")
	} else {
		utils.RespFail(c.Writer, "添加好友失败")
	}
}
