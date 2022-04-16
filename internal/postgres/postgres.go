package postgres

import (
	"context"
	"database/sql"

	"github.com/getpolygon/corexp/internal/gen/postgres_codegen"
	"github.com/getpolygon/corexp/internal/settings"
	_ "github.com/lib/pq"
)

const driverName string = "postgres"

// This function will return a new PostgreSQL connection, which will be
// provided as a dependency by the Fx library if nothing fails.
func New(s *settings.Settings) (*postgres_codegen.Queries, error) {
	// Opening a new connection via the driver provided by the
	// `pq` library.
	conn, err := sql.Open(driverName, s.Postgres)
	if err != nil {
		return nil, err
	}

	// Sending a ping request, to be assured that the connection
	// has been established successfully.
	err = conn.PingContext(context.Background())
	if err != nil {
		return nil, err
	}

	return postgres_codegen.New(conn), nil
}
