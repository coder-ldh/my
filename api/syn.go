package api

import (
	"github.com/gin-gonic/gin"
	"my/global/response"
	"my/model"
)

// @Description mysql数据同步到ES
// @Accept  json
// @Produce  json
// @Success 200 {object} response.Response
// @Header 200 {string} x-token "qwerty"
// @Failure 500 {object} response.Response
// @Router /syn [get]
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
