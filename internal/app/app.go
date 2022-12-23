package app

import (
	"log"

	httpControl "github.com/glhfuck/turbo-waffle/internal/controller/http"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository"
	"github.com/glhfuck/turbo-waffle/internal/usecase"
	"github.com/glhfuck/turbo-waffle/pkg/httpserver"
)

func Run(port string) {
	repository := repository.NewRepository()
	usecase := usecase.NewUsecase(repository)
	controller := httpControl.NewController(usecase)

	httpServer := new(httpserver.Server)
	err := httpServer.Run(port, httpControl.NewRouter(controller))
	if err != nil {
		log.Fatalf("Can not run http server: %s", err.Error())
	}
}
