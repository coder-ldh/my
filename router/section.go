package router

import (
	"github.com/gin-gonic/gin"
	"my/api"
)

func InitSectionRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("section")
	//.Use(middleware.JWTAuth())
	{
		ApiRouter.GET("/list", api.Sections)
		ApiRouter.GET("/book/:bookId", api.SectionListByBookId)
	}
}
