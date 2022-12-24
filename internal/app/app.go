package app

import (
	"log"

	httpControl "github.com/glhfuck/turbo-waffle/internal/controller/http"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository"
	"github.com/glhfuck/turbo-waffle/internal/usecase"
	"github.com/glhfuck/turbo-waffle/pkg/httpserver"
)

func Run() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "9999",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatalf("Can not initialize db: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repository)
	controller := httpControl.NewController(usecase)

	httpServer := new(httpserver.Server)
	err = httpServer.Run(port, httpControl.NewRouter(controller))
	if err != nil {
		log.Fatalf("Can not run http server: %s", err.Error())
	}
}
