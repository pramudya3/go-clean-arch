package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pramudya3/go-clean-arch/domain"
)

func JwtAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		t := strings.Split(auth, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, _ := domain.IsTokenAuthorized(authToken, secret)
			if !authorized {
				c.JSON(http.StatusUnauthorized, domain.ResponseFailed("User Unathorized"))
				c.Abort()
				return
			}
			userID, _ := domain.ExtractIDToken(authToken, secret)
			c.Set("x-user-id", userID)
			c.Next()
		}
	}
}
