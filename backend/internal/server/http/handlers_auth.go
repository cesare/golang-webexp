package http

import (
	"net/http"
	"webexp/internal/auth"
	"webexp/internal/configs"

	"github.com/gin-gonic/gin"
)

func CreateAuthRoutes(config *configs.Config, group *gin.RouterGroup) {
	group.GET("", func(c *gin.Context) {
		authAttrs := auth.NewAuthStart(config).Execute()
		c.JSON(http.StatusOK, authAttrs)
	})
}
