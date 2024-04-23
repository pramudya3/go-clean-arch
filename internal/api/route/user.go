package route

import (
	"github.com/gin-gonic/gin"
	"github.com/pramudya3/go-clean-arch/domain"
	ctrl "github.com/pramudya3/go-clean-arch/internal/api/controller/user"
	"github.com/pramudya3/go-clean-arch/internal/api/middleware"
	"github.com/pramudya3/go-clean-arch/internal/repository/user"
	uc "github.com/pramudya3/go-clean-arch/internal/usecase/user"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRoute(env *domain.Env, db *mongo.Database, g *gin.RouterGroup) {
	userPath := "/users"

	ur := user.NewUserRepository(db)
	userUC := uc.NewUserUsecase(ur, env)

	// public route
	group := g.Group(userPath)
	group.POST("/signup", ctrl.Signup(userUC))
	// add: other public route

	// protected route
	group.Use(middleware.JwtAuth(env.SecretJWT))
	group.GET("/profile", ctrl.GetProfile(userUC))
	group.POST("/signout", ctrl.Signout(userUC))
	// add other protected route
}
