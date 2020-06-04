package api

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"my/global/response"
	"my/service"
)

type LoginParam struct {
	LoginName string
	Password  string
}

// @Tags admin
// @Description 登陆
// @Summary 登陆
// @accept application/json
// @Produce application/json
// @Param data body LoginParam true "{"loginName": "admin","password": "123456"}"
// @Success 200 {object} response.Response
// @Header 200 {string} x-token "qwerty"
// @Failure 500 {object} response.Response
// @Router /admin/login [post]
func Login(c *gin.Context) {
	var login LoginParam
	_ = c.ShouldBindJSON(&login)
	name, err := service.GetAdminByLoginName(login.LoginName)
	if err != nil {
		response.FailMsg(c, "用户名不存在")
		return
	}
	m := md5.New()
	m.Write([]byte(login.Password))
	if hex.EncodeToString(m.Sum(nil)) == *name.Password {
		response.SuccessObj(c, name)
		return
	}
	response.FailMsg(c, "密码有误")
	return
}
