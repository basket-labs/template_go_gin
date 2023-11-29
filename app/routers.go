package app

import (
	"github.com/gin-gonic/gin"
)

func healthz(c *gin.Context) {
	c.Data(200, gin.MIMEJSON, []byte("OK"))
}

func version(c *gin.Context) {
	c.Data(200, gin.MIMEJSON, []byte(GetVersion()))
}
