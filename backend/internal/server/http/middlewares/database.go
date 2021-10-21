package middlewares

import (
	"webexp/internal/configs"

	"github.com/gin-gonic/gin"
)

func DatabaseMiddleware(config *configs.Config) (gin.HandlerFunc, error) {
	handler := func(c *gin.Context) {
	}
	return handler, nil
}
