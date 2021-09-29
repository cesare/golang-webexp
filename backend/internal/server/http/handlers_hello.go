package http

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	session := sessions.Default(c)

	name := c.Query("name")
	if name != "" {
		session.Set("last-visited", name)
		session.Save()
	} else {
		name = session.Get("last-visited").(string)
	}

	message := fmt.Sprintf("hello, %s", name)
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
