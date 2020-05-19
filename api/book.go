package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my/models"
	"net/http"
)

func Books(c *gin.Context) {

	result, err := models.Books()
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到相关信息",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": result,
	})
}

func BookSectionByNum(c *gin.Context) {
	num := c.Param("num")
	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": num,
	})
}
