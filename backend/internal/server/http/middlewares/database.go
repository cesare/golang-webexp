package middlewares

import (
	"database/sql"
	"webexp/internal/configs"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func DatabaseMiddleware(config *configs.Config) (gin.HandlerFunc, error) {
	db, err := sql.Open("postgres", config.Database.ConnectionString())
	if err != nil {
		return nil, err
	}

	handler := func(c *gin.Context) {
		c.Set("database", db)
		c.Next()
	}
	return handler, nil
}
