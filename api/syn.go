package api

import (
	"github.com/gin-gonic/gin"
	"my/global/response"
	"my/model"
)

func Syn(c *gin.Context) {
	err := model.BookMysqlToEs()
	error := model.SectionMysqlToEs()
	if err != nil {
		response.FailMsg(c, err.Error())
		return
	}
	if error != nil {
		response.FailMsg(c, error.Error())
		return
	}
	response.Success(c)
}
