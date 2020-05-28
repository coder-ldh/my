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
		ApiRouter.GET("/list/:bookId", api.SectionListByBookId)
		ApiRouter.GET("/one/:sectionId", api.SectionById)
	}
}
