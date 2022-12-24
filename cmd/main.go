package main

import (
	"log"
	"os"

	"github.com/glhfuck/turbo-waffle/internal/app"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository"
	"github.com/glhfuck/turbo-waffle/pkg/httpserver"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	err := InitConfig()
	if err != nil {
		log.Fatalf("Can't initialize config: %s", err.Error())
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Can't load env variables: %s", err.Error())
	}

	cfg := app.Config{
		HttpConfig: httpserver.HttpConfig{
			Port: viper.GetString("port"),
		},
		PostgresConfig: repository.PostgresConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
		},
	}

	app.Run(cfg)
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
