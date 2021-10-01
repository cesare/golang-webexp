package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAuthRoutes(group *gin.RouterGroup) {
	group.GET("", authStart)
}

func authStart(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "dummy",
	})
}
