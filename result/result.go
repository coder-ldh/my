package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 501,
		"msg":  msg,
	})
}

func Success(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  msg,
	})
}

func SuccessObj(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": obj,
	})
}
