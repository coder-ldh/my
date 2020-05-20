package api

import (
	"github.com/gin-gonic/gin"
	"my/models"
	"my/result"
)

func Syn(c *gin.Context) {
	err := models.BookMysqlToEs()
	error := models.SectionMysqlToEs()
	if err != nil {
		result.Fail(c, err.Error())
		return
	}
	if error != nil {
		result.Fail(c, error.Error())
		return
	}
	result.Success(c, "操作成功")
}
