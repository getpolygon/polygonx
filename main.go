package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"polygon.am/core/pkg/config"
	"polygon.am/core/pkg/types"
)

// Global, configuration variable for accessing and changing
// the configuration on demand.
var Configuration types.Config

// The default path for looking for the default configuration
// file path, if the environment variable was not supplied.
const DefaultConfigurationFilePath string = "./.conf.yaml"

func init() {
	path, err := filepath.Abs(DefaultConfigurationFilePath)
	if err != nil {
		log.Fatal(err)
	}

	config, err := config.ParseConfig(path)
	if err != nil {
		log.Fatal(err)
	}

	// Assigning parsed configuration to a global variable
	Configuration = *config

}

func main() {
	router := chi.NewRouter()

	// Binding the middleware to chi router
	router.Use(middleware.Logger)
	router.Use(middleware.GetHead)
	router.Use(middleware.NoCache)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/status"))
	// This middleware will handle any internal crashes inside
	// of our application and will send an `500: Internal Server Error`
	// response.
	router.Use(middleware.Recoverer)

	log.Println("getpolygon/corexp started at http://" + Configuration.Polygon.Addr)

	// Binding to the address specified or defaulted to from the configuration
	// and attaching chi routes to the server.
	http.ListenAndServe(Configuration.Polygon.Addr, router)
}
