package http

import (
	"net/http"
	"webexp/internal/auth"
	"webexp/internal/configs"

	"github.com/gin-gonic/gin"
)

func CreateAuthRoutes(config *configs.Config, group *gin.RouterGroup) {
	group.GET("", func(c *gin.Context) {
		authAttrs, err := auth.NewAuthStart(config).Execute()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"state":            authAttrs.State,
			"authorizationUri": authAttrs.AuthorizationUri,
		})
	})
}
