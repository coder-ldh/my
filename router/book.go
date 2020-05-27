package router

import (
	"github.com/gin-gonic/gin"
	"my/api"
)

func InitBookRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("book")
	{
		ApiRouter.GET("/list", api.Books)
		ApiRouter.GET("/query", api.Query)
		ApiRouter.GET("/detail/:bookId", api.BookById)
	}
}
