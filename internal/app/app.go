package app

import (
	"github.com/sirupsen/logrus"

	"github.com/glhfuck/turbo-waffle/internal/controller/httpcontroller"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository/pgrepository"
	"github.com/glhfuck/turbo-waffle/internal/usecase"
	"github.com/glhfuck/turbo-waffle/pkg/httpserver"
)

func Run(config *Config) {
	db, err := pgrepository.NewPostgresDB(config.DBConfig)

	if err != nil {
		logrus.Fatalf("Can't initialize db: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repository)
	controller := httpcontroller.NewController(usecase)

	httpServer := new(httpserver.Server)
	err = httpServer.Run(config.HttpConfig, httpcontroller.NewRouter(controller))
	if err != nil {
		logrus.Fatalf("Can't run http server: %s", err.Error())
	}
}
