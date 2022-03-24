package persistence

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

// PostgreSQL connection that will be used for making queries
// without creating additional clients.
var Postgres pgx.Conn

func init() {
	dsn := os.Getenv("POLYGON_CORE_DATABASES_POSTGRES")
	// Connecting to PostgreSQL using the provided connection URL
	// from the environment variable.
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Trying to ping the database, to verify that the connection
	// was initialized successfully.
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// If all the procedures executed successfully, assigning
	// the connection to a global variable.
	Postgres = *conn

	log.Println("connected to postgresql")
}
