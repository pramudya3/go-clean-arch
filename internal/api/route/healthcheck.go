package route

import (
	"github.com/gin-gonic/gin"
	"github.com/pramudya3/go-clean-arch/internal/api/controller/healthcheck"
)

func NewHealthchek(g *gin.RouterGroup) {
	g = g.Group("")
	g.GET("/healthcheck", healthcheck.Healthcheck())
}
