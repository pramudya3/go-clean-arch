package domain

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv             string `mapstructure:"APP_ENV"`
	ServerAddress      string `mapstructure:"SERVER_ADDRESS"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBUser             string `mapstructure:"DB_USER"`
	DBPass             string `mapstructure:"DB_PASS"`
	DBName             string `mapstructure:"DB_NAME"`
	SecretJWT          string `mapstructure:"SECRET_JWT"`
	ContextTimeout     int    `mapstructure:"CONTEXT_TIMEOUT"`
	ExpiryAccessToken  int    `mapstructure:"EXPIRY_ACCESS_TOKEN"`
	ExpiryRefreshToken int    `mapstructure:"EXPIRY_REFRESH_TOKEN"`
}

func InitEnv(filepath string) (*Env, error) {
	env := Env{}

	viper.SetConfigFile(filepath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("can't find the env file: %v", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		return nil, fmt.Errorf("failed to unmarshalling env: %v", err)
	}

	if env.AppEnv == "development" {
		log.Println("this app is running in development env")
	}

	return &env, nil
}
