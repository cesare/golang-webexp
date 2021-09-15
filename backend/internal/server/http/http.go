package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Engine() *gin.Engine {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	return engine
}
