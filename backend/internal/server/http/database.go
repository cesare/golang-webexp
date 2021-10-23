package http

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func database(c *gin.Context) *sql.DB {
	v := c.MustGet("database")
	return v.(*sql.DB)
}
