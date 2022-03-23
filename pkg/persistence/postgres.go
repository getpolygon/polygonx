package persistence

import (
	"context"
	"database/sql"
	"log"
	"os"
)

// A global variable for accessing PostgreSQL connection pool
// and using it in various operations.
var Postgres *sql.DB

func init() {
	// Getting PostgreSQL connection string from an environment variable,
	// and failing automatically, by connecting, if the url was not to an
	// invalid address.
	//
	// TODO: Update the name for PostgreSQL connection string environment variable
	dsn := os.Getenv("")

	conn, connerr := sql.Open("pg", dsn)
	if connerr != nil {
		log.Fatal(connerr)
	}

	// Sending a simple test query to verify connection validity, and assign
	// the global database connection variable.
	pingerr := conn.PingContext(context.Background())
	if pingerr != nil {
		log.Fatal(pingerr)
	}

	// Assinging the connection to a global variable if all the procedures
	// have been executed successfully.
	Postgres = conn
}
