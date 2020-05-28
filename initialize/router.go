package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"my/api"
	_ "my/docs"
	"my/global"
	"my/middleware"
	"my/router"
)

//初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	//Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Debug("use middleware logger")
	/*跨域*/
	Router.Use(middleware.Cors())
	global.GVA_LOG.Debug("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.POST("/syn", api.Syn)
	global.GVA_LOG.Debug("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	router.InitBookRouter(ApiGroup)
	router.InitSectionRouter(ApiGroup)
	router.InitCaptchaRouter(ApiGroup)
	global.GVA_LOG.Info("router register success")
	return Router
}
