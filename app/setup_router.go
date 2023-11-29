package app

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r gin.IRoutes) {
	r.GET("/healthz", healthz)
	r.GET("/version", version)
}
