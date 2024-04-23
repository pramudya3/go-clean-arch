package route

import (
	"github.com/gin-gonic/gin"
	"github.com/pramudya3/go-clean-arch/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(env *domain.Env, db *mongo.Database, g *gin.Engine) {
	group := g.Group("/api/v1")
	NewUserRoute(env, db, group)
	NewHealthchek(group)
}
