package core

import (
	"fmt"
	"my/global"
	"my/initialize"
)

func RunWindowsServer() {
	initialize.Redis()
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")
	fmt.Printf(`欢迎使用 Gin-Vue-Admin
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
`, address)
	Router.Run(address)
}
