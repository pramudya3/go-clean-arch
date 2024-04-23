package config

import (
	"context"
	"log"

	"github.com/pramudya3/go-clean-arch/internal/api/route"
	"github.com/pramudya3/go-clean-arch/pkg/mongodb"
)

func (a *initApp) Run() {
	db, err := mongodb.InitMongoDB(a.Env)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Client().Disconnect(context.TODO())

	route.SetupRouter(a.Env, db, a.Gin)

	a.NewServer()
}
