package api

import (
	"github.com/gin-gonic/gin"
	"my/models"
	"net/http"
)

func Syn(c *gin.Context) {
	err := models.BookMysqlToEs()
	error := models.SectionMysqlToEs()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "500",
			"msg":  err,
		})
		return
	}
	if error != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": "500",
			"msg":  error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"msg":  "操作成功",
	})
}
