package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/glhfuck/turbo-waffle/internal/controller/httpcontroller"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository"
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository/pgrepository"
	"github.com/glhfuck/turbo-waffle/internal/usecase"
	"github.com/glhfuck/turbo-waffle/pkg/httpserver"
)

func Run(config *Config) {
	logrus.Println("Connecting to the database")
	db, err := pgrepository.NewPostgresDB(config.DBConfig)

	if err != nil {
		logrus.Fatalf("Can't initialize db: %s", err.Error())
	}

	logrus.Println("Successfully connected to the database")

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repository)
	controller := httpcontroller.NewController(usecase)

	httpServer := new(httpserver.Server)

	go func() {
		err = httpServer.Run(config.HttpConfig, httpcontroller.NewRouter(controller))
		if err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Can't run http server: %s", err.Error())
		}
	}()

	logrus.Println("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Server is shuting down")

	err = httpServer.Shutdown(context.Background())
	if err != nil {
		logrus.Errorf("Error occured on server shutting down: %s", err)
	}

	err = db.Close()
	if err != nil {
		logrus.Errorf("Error occured on database connection close: %s", err)
	}

	logrus.Println("Server shut down successfully")
}
