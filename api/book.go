package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"log"
	"my/global"
	"my/global/response"
	"my/model"
	"strconv"
)

type Document struct {
	Id         int    `json:"id"`
	BookName   string `json:"bookName"`
	BookIntro  string `json:"bookIntro"`
	BookAuthor string `json:"bookAuthor"`
}

type DocumentResponse struct {
	BookName   string `json:"bookName"`
	BookIntro  string `json:"bookIntro"`
	BookAuthor string `json:"bookAuthor"`
}
type SearchResponse struct {
	Time      string             `json:"time"`
	Hits      string             `json:"hits"`
	Documents []DocumentResponse `json:"documents"`
}

// @Description 获取书籍列表
// @Accept  json
// @Produce  json
// @Param pageNum query int true "1"
// @Param pageSize query int true "10"
// @Success 200 {object} response.Response
// @Header 200 {string} x-token "qwerty"
// @Failure 500 {object} response.Response
// @Router /book/list [get]
func Books(c *gin.Context) {
	pageNum := 1
	pageSize := 10
	if i, err := strconv.Atoi(c.Query("pageNum")); err == nil {
		pageNum = i
	}
	if i, err := strconv.Atoi(c.Query("pageSize")); err == nil {
		pageSize = i
	}
	results, err := model.Books(pageNum, pageSize)
	if err != nil {
		fmt.Print(err)
		response.FailMsg(c, err.Error())
		return
	}
	response.SuccessObj(c, results)
}

// @Description 搜索书
// @Accept  json
// @Produce  json
// @Param query query string true "三体"
// @Param pageNum query int true "1"
// @Param pageSize query int true "10"
// @Success 200 {object} response.Response
// @Header 200 {string} x-token "qwerty"
// @Failure 500 {object} response.Response
// @Router /book/query [get]
func Query(c *gin.Context) {
	key := c.Query("query")
	if key == "" {
		response.FailMsg(c, "Query not specified")
		return
	}
	pageNum := 1
	pageSize := 10
	if i, err := strconv.Atoi(c.Query("pageNum")); err == nil {
		pageNum = i
	}
	if i, err := strconv.Atoi(c.Query("pageSize")); err == nil {
		pageSize = i
	}
	esQuery := elastic.NewMultiMatchQuery(key, "BookName", "BookIntro", "BookAuthor").
		Fuzziness("2")
	searchResult, err := global.GVA_ES.Search().Index("book").Query(esQuery).
		From(pageNum - 1).Size(pageSize).
		Do(c.Request.Context())
	if err != nil {
		log.Println(err)
		response.FailMsg(c, "Something went wrong")
		return
	}
	res := SearchResponse{
		Time: fmt.Sprintf("%d", searchResult.TookInMillis),
		Hits: fmt.Sprintf("%d", searchResult.Hits.TotalHits),
	}
	docs := make([]DocumentResponse, 0)
	for _, hit := range searchResult.Hits.Hits {
		var doc DocumentResponse
		json.Unmarshal(hit.Source, &doc)
		docs = append(docs, doc)
	}
	res.Documents = docs
	response.SuccessObj(c, res)
	return
}

// @Description 书详情
// @Accept  json
// @Produce  json
// @Param bookId path int true "1"
// @Success 200 {object} response.Response
// @Header 200 {string} x-token "qwerty"
// @Failure 500 {object} response.Response
// @Router /book/detail/{bookId} [get]
func BookById(c *gin.Context) {
	bookId := c.Param("bookId")
	if bookId == "" {
		response.FailMsg(c, "bookId not specified")
		return
	}
	book, err := model.GetBookByIdFromEs(bookId)
	if err != nil {
		response.FailMsg(c, err.Error())
	}
	response.SuccessObj(c, book)
}
