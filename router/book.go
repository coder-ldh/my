package router

import (
	"github.com/gin-gonic/gin"
	"my/api"
	"my/middleware"
)

func InitBookRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("book").Use(middleware.JWTAuth())
	//.Use(middleware.CasbinHandler())
	{
		ApiRouter.GET("/list", api.Books)
		ApiRouter.GET("/query", api.BookQuery)
		ApiRouter.GET("/detail/:bookId", api.BookById)
		ApiRouter.GET("/section/:num", api.BookSectionByNum)
	}
}
