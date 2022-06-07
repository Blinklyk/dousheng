package main

import (
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/initialize"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := initialize.Init(); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	initRouter(r)
	r.Run(":" + global.App.Config.App.Port)
}
