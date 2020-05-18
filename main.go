package main

import (
	"my/router"
)

func main() {
	r := router.InnitRouter()

	r.Run(":9090")
}
