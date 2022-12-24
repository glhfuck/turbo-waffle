package app

import (
	postgresrepo "github.com/glhfuck/turbo-waffle/internal/infrastructure/repository/postgres_repo"
	"github.com/glhfuck/turbo-waffle/pkg/httpserver"
)

type Config struct {
	httpserver.HttpConfig
	postgresrepo.PostgresConfig
}
