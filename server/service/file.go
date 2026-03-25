package service

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"server/define"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("fileResource")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "上传文件失败",
		})
		return
	}
	fmt.Println(file.Filename)
	fmt.Println(file.Size)
	fileExt := path.Ext(file.Filename)
	fmt.Printf("文件后缀：%s", fileExt)
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" && fileExt != ".gif" && fileExt != ".bmp" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "上传文件格式错误",
		})
		return
	}
	fileNameInt := time.Now().Unix()
	fileNameStr := strconv.FormatInt(fileNameInt, 10)
	fileName := fileNameStr + fileExt
	filePath := filepath.Join(define.StaticResource, "/", fileName)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "上传文件失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"msg":      "上传成功",
		"fileName": fileName,
	})
}
