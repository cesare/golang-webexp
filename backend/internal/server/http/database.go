package http

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func database(c *gin.Context) (*sql.DB, error) {
	v, exists := c.Get("database")
	if !exists {
		return nil, fmt.Errorf("database missing")
	}

	db, ok := v.(*sql.DB)
	if !ok {
		return nil, fmt.Errorf("unknown database object")
	}

	return db, nil
}
