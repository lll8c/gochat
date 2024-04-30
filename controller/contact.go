package controller

import (
	"fmt"
	"gochat/models"
	"gochat/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SearchFriends 进入chat主页时，查询所有好友
func SearchFriends(c *gin.Context) {
	str := c.PostForm("userId")
	userID, _ := strconv.Atoi(str)
	fmt.Println(userID)
	users := models.SearchFriend(uint(userID))
	utils.RespOKList(c.Writer, users, len(users))
}
