package controller

import (
	"fmt"
	"gochat/utils"
	"io"
	"math/rand/v2"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// Upload 上传文件到本地
func Upload(c *gin.Context) {
	srcFile, head, err := c.Request.FormFile("file")
	if err != nil {
		utils.RespFail(c.Writer, err.Error())
		return
	}
	suffix := ".png"
	ofilName := head.Filename
	tem := strings.Split(ofilName, ".")
	if len(tem) > 1 {
		suffix = "." + tem[len(tem)-1]
	}
	fileName := fmt.Sprintf("%d%04d%s", time.Now().Unix(), rand.Int32(), suffix)
	dstFile, err := os.Create("./asset/upload/" + fileName)
	if err != nil {
		utils.RespFail(c.Writer, err.Error())
		return
	}
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		utils.RespFail(c.Writer, err.Error())
		return
	}
	//
	url := "./asset/upload/" + fileName
	utils.RespOK(c.Writer, "发送图片成功", url)
}
