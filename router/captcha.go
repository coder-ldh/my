package router

import (
	"github.com/gin-gonic/gin"
	"my/api"
)

func InitCaptchaRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("base")
	{
		ApiRouter.GET("/captcha", api.Captcha)
		ApiRouter.GET("/captcha/:captchaId", api.CaptchaImg)
	}
}
