package middleware

import (
	"github.com/gin-gonic/gin"
	"my/global"
	"my/global/response"
	"my/model/request"
	"my/service"
)

/*拦截器*/
func InterceptorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		waitUse := claims.(*request.CustomClaims)
		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := waitUse.AuthorityId
		e := service.Casbin()
		//判断策略中是否存在
		if global.GVA_CONFIG.System.Env == "develop" || e.Enforce(sub, obj, act) {
			c.Next()
		} else {
			response.Result(response.ERROR, gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
	}
}
