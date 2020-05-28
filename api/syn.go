package api

import (
	"github.com/gin-gonic/gin"
	"my/global/response"
	"my/service"
)

// @Tags base
// @Summary mysql数据同步到ES
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response
// @Header 200 {string} x-token "qwerty"
// @Failure 500 {object} response.Response
// @Router /syn [post]
func Syn(c *gin.Context) {
	err := service.BookMysqlToEs()
	error := service.SectionMysqlToEs()
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
