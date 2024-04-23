package config

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pramudya3/go-clean-arch/domain"
)

type initApp struct {
	Env *domain.Env
	Gin *gin.Engine
}

func NewConfig() *initApp {
	env, err := domain.InitEnv("./../.env")
	if err != nil {
		env, err = domain.InitEnv(".env")
		if err != nil {
			log.Fatal(err)
		}
	}

	if env.AppEnv != "development" {
		gin.SetMode(gin.ReleaseMode)
		log.Print("production mode\n")
	}
	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())

	return &initApp{
		Env: env,
		Gin: ginEngine,
	}
}
