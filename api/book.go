package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"log"
	orm "my/database"
	"my/models"
	"my/result"
	"net/http"
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

func BookQuery(c *gin.Context) {
	key := c.Query("query")
	if key == "" {
		result.Fail(c, "Query not specified")
		return
	}
	pageNum := 0
	pageSize := 10
	if i, err := strconv.Atoi(c.Query("pageNum")); err == nil {
		pageNum = i
	}
	if i, err := strconv.Atoi(c.Query("pageSize")); err == nil {
		pageSize = i
	}
	esQuery := elastic.NewMultiMatchQuery(key, "BookName", "BookIntro", "BookAuthor").
		Fuzziness("2")
	searchResult, err := orm.Es.Search().Index("book").Query(esQuery).
		From(pageNum).Size(pageSize).
		Do(c.Request.Context())
	if err != nil {
		log.Println(err)
		result.Fail(c, "Something went wrong")
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
	result.SuccessObj(c, res)
	return
}

func BookSectionByNum(c *gin.Context) {
	num := c.Param("num")
	c.JSON(http.StatusOK, gin.H{
		"code":    -1,
		"message": num,
	})
}
