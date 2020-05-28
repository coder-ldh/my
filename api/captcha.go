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
// @Router /base/captcha [get]
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

// @Tags base
// @Summary 校验验证吗
// @accept application/json
// @Produce application/json
// @Param captchaId query string true "1"
// @Param Captcha query string true "1"
// @Success 200 {string} string "{"Code":200,"data":{},"msg":"获取成功"}"
// @Router /base/captcha/verify [post]
func VerifyCode(c *gin.Context) {
	captchaId := c.Query("captchaId")
	if captchaId == "" {
		response.FailMsg(c, "captchaId not specified")
		return
	}
	Captcha := c.Query("Captcha")
	if Captcha == "" {
		response.FailMsg(c, "Captcha not specified")
		return
	}
	if utils.VerifyCode(captchaId, Captcha) {
		response.SuccessMsg(c, "验证成功")
	} else {
		response.FailMsg(c, "验证码有误")
	}
}
