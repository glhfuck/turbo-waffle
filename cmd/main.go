package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/glhfuck/turbo-waffle/internal/app"
	postgresrepo "github.com/glhfuck/turbo-waffle/internal/infrastructure/repository/postgres_repo"
	"github.com/glhfuck/turbo-waffle/pkg/httpserver"
)

func main() {
	err := InitConfig()
	if err != nil {
		logrus.Fatalf("Can't initialize config: %s", err.Error())
	}

	err = godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Can't load env variables: %s", err.Error())
	}

	cfg := app.Config{
		HttpConfig: httpserver.HttpConfig{
			Port: viper.GetString("port"),
		},
		PostgresConfig: postgresrepo.PostgresConfig{
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
