package app

import (
	"github.com/sirupsen/logrus"

	httpControl "github.com/glhfuck/turbo-waffle/internal/controller/httpcontroller"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository/pgrepository"
	"github.com/glhfuck/turbo-waffle/internal/usecase"
	"github.com/glhfuck/turbo-waffle/pkg/httpserver"
)

func Run(cfg Config) {
	db, err := pgrepository.NewPostgresDB(cfg.PostgresConfig)

	if err != nil {
		logrus.Fatalf("Can't initialize db: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repository)
	controller := httpControl.NewController(usecase)

	httpServer := new(httpserver.Server)
	err = httpServer.Run(cfg.HttpConfig, httpControl.NewRouter(controller))
	if err != nil {
		logrus.Fatalf("Can't run http server: %s", err.Error())
	}
}
