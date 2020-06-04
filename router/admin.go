package router

import (
	"github.com/gin-gonic/gin"
	"my/api"
)

func InitAdminRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("admin")
	{
		ApiRouter.POST("/login", api.Login)
	}
}
