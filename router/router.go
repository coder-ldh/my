package router

import (
	"github.com/gin-gonic/gin"
	. "my/api"
)

func InnitRouter() *gin.Engine {
	Engine := gin.Default()
	Engine.GET("/book/list", Books)
	Engine.GET("/book/query", Books)
	Engine.GET("/book/section/:num", BookSectionByNum)
	Engine.GET("/section/list", Books)
	Engine.GET("/section/{}", Books)
	return Engine
}
