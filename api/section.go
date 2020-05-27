package api

import (
	"github.com/gin-gonic/gin"
	"my/global/response"
	"my/model"
	"strconv"
)

func Sections(c *gin.Context) {
	pageNum := 1
	pageSize := 10
	if i, err := strconv.Atoi(c.Query("pageNum")); err == nil {
		pageNum = i
	}
	if i, err := strconv.Atoi(c.Query("pageSize")); err == nil {
		pageSize = i
	}
	sections, err := model.Sections(pageNum, pageSize)
	if err != nil {
		response.FailMsg(c, err.Error())
	}
	response.SuccessObj(c, sections)
}
