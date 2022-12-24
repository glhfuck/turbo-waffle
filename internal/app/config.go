package app

import (
	"github.com/glhfuck/turbo-waffle/internal/infrastructure/repository"
	"github.com/glhfuck/turbo-waffle/pkg/httpserver"
)

type Config struct {
	httpserver.HttpConfig
	repository.PostgresConfig
}
