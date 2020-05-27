package main

import (
	"my/core"
	"my/initialize"
)

func main() {
	initialize.Mysql()
	core.RunWindowsServer()
}
