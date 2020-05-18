package router
import (
	"github.com/gin-gonic/gin"
	."my/api"
)
func InnitRouter() *gin.Engine{
	Engine := gin.Default()
	Engine.GET("/user/list", Books)
	return Engine
}