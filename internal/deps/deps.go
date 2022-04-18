package deps

import (
	"github.com/getpolygon/corexp/internal/gen/postgres_codegen"
	"github.com/getpolygon/corexp/internal/postgres"
	"github.com/getpolygon/corexp/internal/settings"
)

type Dependencies struct {
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

	return &Dependencies{
		Postgres: postgres,
		Settings: settings,
	}, nil
}
