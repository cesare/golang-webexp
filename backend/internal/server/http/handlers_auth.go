package http

import (
	"net/http"
	"webexp/internal/auth"
	"webexp/internal/auth/identity"
	"webexp/internal/webexp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateAuthRoutes(context *webexp.Context, group *gin.RouterGroup) {
	group.POST("", func(c *gin.Context) {
		authAttrs, err := auth.NewAuthStart(context.Config()).Execute()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		session := sessions.Default(c)
		session.Clear()
		session.Set("auth-state", authAttrs.State)
		session.Save()

		c.JSON(http.StatusOK, gin.H{
			"authorizationUri": authAttrs.AuthorizationUri,
		})
	})

	type calbackRequest struct {
		State string `json:"state" binding:"required"`
		Code  string `json:"code"  binding:"required"`
	}

	group.POST("/callback", func(c *gin.Context) {
		var request calbackRequest
		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid request",
			})
			return
		}

		session := sessions.Default(c)
		savedState, ok := session.Get("auth-state").(string)
		if !ok || savedState != request.State {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "states not match",
			})
			return
		}

		attrs := auth.CallbackAttributes{
			Code:  request.Code,
			State: request.State,
		}
		results, err := auth.NewAuth(context.Config(), attrs).Execute()
		if err != nil {
			switch e := err.(type) {
			case *auth.AuthRejected:
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": e.Error(),
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": e.Error(),
				})
			}
			return
		}

		registration := identity.NewIdentityRegistration(context.Database(), results.Identifier)
		identity, err := registration.Execute()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		session.Clear()
		session.Set("sub", identity.Id)
		session.Save()

		c.JSON(http.StatusCreated, gin.H{
			"sub": identity.Id,
		})
	})
}
