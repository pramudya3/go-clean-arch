package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/pramudya3/go-clean-arch/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongoDB(env *domain.Env) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := ""
	if env.DBUser == "" || env.DBPass == "" {
		uri = fmt.Sprintf("mongodb://%s:%s", env.DBHost, env.DBPort)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s", env.DBUser, env.DBPass, env.DBHost, env.DBPort)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("connect database failed, err: %v", err.Error())
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("ping client failed, err: %v", err)
	}

	return client.Database(env.DBName), nil
}
