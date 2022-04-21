package deps

import (
	"context"

	"github.com/getpolygon/corexp/internal/gen/postgres_codegen"
	"github.com/getpolygon/corexp/internal/services/postgres"
	"github.com/getpolygon/corexp/internal/settings"
	"github.com/go-redis/redis/v8"
)

type Dependencies struct {
	Redis    *redis.Client
	Settings *settings.Settings
	Postgres *postgres_codegen.Queries
}

func New() (*Dependencies, error) {
	settings, err := settings.New()
	if err != nil {
		return nil, err
	}

	postgres, err := postgres.New(settings)
	if err != nil {
		return nil, err
	}

	redis := redis.NewClient(&redis.Options{
		Addr: settings.Redis,
	})
	if err := redis.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &Dependencies{
		Redis:    redis,
		Postgres: postgres,
		Settings: settings,
	}, nil
}
