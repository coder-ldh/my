package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = http.StatusBadRequest
	SUCCESS = http.StatusOK
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func Fail(c *gin.Context, msg string) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func Success(c *gin.Context, msg string) {
	Result(SUCCESS, map[string]interface{}{}, msg, c)
}

func SuccessObj(c *gin.Context, obj interface{}) {
	Result(SUCCESS, obj, "操作成功", c)
}
