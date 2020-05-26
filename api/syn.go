package api

import (
	"github.com/gin-gonic/gin"
	"my/global/response"
	"my/models"
)

func Syn(c *gin.Context) {
	err := models.BookMysqlToEs()
	error := models.SectionMysqlToEs()
	if err != nil {
		response.Fail(c, err.Error())
		return
	}
	if error != nil {
		response.Fail(c, error.Error())
		return
	}
	response.Success(c, "操作成功")
}
