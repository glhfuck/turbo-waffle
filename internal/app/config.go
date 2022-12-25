package app

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository/pgrepository"
	"github.com/glhfuck/turbo-waffle/pkg/httpserver"
)

type Config struct {
	httpserver.HttpConfig
	pgrepository.DBConfig
}

func NewConfig() *Config {
	if err := InitConfig(); err != nil {
		logrus.Fatalf("Can't initialize configs: %s", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatalf("Can't load env variables: %s", err.Error())
	}

	config := &Config{
		httpserver.HttpConfig{
			Port: viper.GetString("http.port"),
		},
		pgrepository.DBConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
		},
	}

	return config
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
