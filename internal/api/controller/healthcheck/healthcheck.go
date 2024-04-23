package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthcheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}
