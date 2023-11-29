package main

import (
	"awesomeProject/app"
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadAppConfig()
	r := gin.Default()

	var api = r.Group("/api")
	app.SetupRouter(api)

	r.Run("0.0.0.0:" + config.AppConfig.Server.Port)
}
