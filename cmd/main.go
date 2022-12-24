package main

import (
	"log"

	"github.com/glhfuck/turbo-waffle/internal/app"
	"github.com/spf13/viper"
)

func main() {
	err := InitConfig()
	if err != nil {
		log.Fatalf("Can't initialize config: %s", err.Error())
	}

	app.Run(viper.GetString("port"))
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
