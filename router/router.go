package router

import (
	"github.com/gin-gonic/gin"
	. "my/api"
	"net/http"
)

func InnitRouter() *gin.Engine {
	Engine := gin.Default()
	Engine.Use(cors())
	Engine.GET("/book/list", Books)
	Engine.GET("/book/query", BookQuery)
	Engine.GET("/book/detail/:bookId", BookById)
	Engine.GET("/book/section/:num", BookSectionByNum)
	Engine.GET("/section/list", Sections)

	Engine.GET("/syn", Syn)
	return Engine
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST,PUT,GET,DELETE,OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
