package http

import (
	"github.com/gin-gonic/gin"
)

func Engine() *gin.Engine {
	engine := gin.Default()
	engine.GET("/", hello)
	return engine
}
