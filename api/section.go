package api

import (
	"github.com/gin-gonic/gin"
	"my/global/response"
	"my/model"
	"strconv"
)

// @Description 章节列表
// @Accept  json
// @Produce  json
// @Param pageNum query int true "1"
// @Param pageSize query int true "10"
// @Success 200 {object} response.Response
// @Header 200 {string} x-token "qwerty"
// @Failure 500 {object} response.Response
// @Router /section/list [get]
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

// @Description 查询书下所有章节信息
// @Accept  json
// @Produce  json
// @Param bookId path int true "1"
// @Success 200 {object} response.Response
// @Header 200 {string} x-token "qwerty"
// @Failure 500 {object} response.Response
// @Router /section/book/{bookId} [get]
func SectionListByBookId(c *gin.Context) {
	bookId := c.Param("bookId")
	if bookId == "" {
		response.FailMsg(c, "bookId not specified")
		return
	}
	sections, err := model.SectionListByBookId(bookId, c)
	if err != nil {
		response.FailMsg(c, err.Error())
	}
	response.SuccessObj(c, sections)
}
