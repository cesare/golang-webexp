package http

import (
	"webexp/internal/configs"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Engine(config *configs.Config) *gin.Engine {
	engine := gin.Default()

	store := cookie.NewStore(config.App.SessionKey.Bytes())
	engine.Use(sessions.Sessions("webexp-session", store))

	engine.GET("/", hello)
	return engine
}
