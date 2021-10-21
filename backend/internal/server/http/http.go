package http

import (
	"webexp/internal/configs"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Engine(config *configs.Config) (*gin.Engine, error) {
	engine := gin.Default()

	store := cookie.NewStore(config.App.SessionKey.Bytes())
	engine.Use(sessions.Sessions("webexp-session", store))
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{config.Frontend.BaseUri},
		AllowMethods: []string{"POST"},
		AllowHeaders: []string{
			"Content-Type",
		},
		AllowCredentials: true,
	}))

	authGroup := engine.Group("/auth")
	CreateAuthRoutes(config, authGroup)

	engine.GET("/", hello)
	return engine, nil
}
