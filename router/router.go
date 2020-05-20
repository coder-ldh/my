package router

import (
	"github.com/gin-gonic/gin"
	. "my/api"
)

func InnitRouter() *gin.Engine {
	Engine := gin.Default()
	Engine.GET("/book/list", Books)
	Engine.GET("/book/query", BookQuery)
	Engine.GET("/book/section/:num", BookSectionByNum)
	Engine.GET("/section/list", Sections)

	Engine.GET("/syn", Syn)
	return Engine
}
