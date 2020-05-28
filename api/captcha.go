package api

import (
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"my/global"
	"my/global/response"
	"my/utils"
)

// @Tags base
// @Summary 生成验证码
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"Code":200,"data":{},"msg":"获取成功"}"
// @Router /base/captcha [post]
func Captcha(c *gin.Context) {
	captchaId := captcha.NewLen(global.GVA_CONFIG.Captcha.KeyLong)
	response.Result(response.SUCCESS, gin.H{
		"CaptchaId": captchaId,
		"PicPath":   "/base/captcha/" + captchaId + ".png",
	}, "获取成功", c)
}

// @Tags base
// @Summary 生成验证码图片路径
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"Code":200,"data":{},"msg":"获取成功"}"
// @Router /base/captcha/:captchaId [get]
func CaptchaImg(c *gin.Context) {
	utils.GinCaptchaServeHTTP(c.Writer, c.Request)
}
