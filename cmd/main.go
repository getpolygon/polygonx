package main

import (
	"log"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"polygon.am/core/api/types"
	"polygon.am/core/pkg/config"
)

// Global, configuration variable for accessing and changing
// the configuration on demand.
var Configuration *types.Config

// The default path for looking for the default configuration
// file path, if the environment variable was not supplied.
const DefaultConfigurationFilePath string = "./.conf.yaml"

func init() {
	// Getting the absolute path for the configuration path
	path, err := filepath.Abs(DefaultConfigurationFilePath)
	if err == nil {
		// Parsing the YAML configuration and deserializing it
		config, err := config.ParseConfig(path)
		if err == nil {
			// Assigning the configuration to a global variable
			Configuration = config
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

func main() {
	server := fiber.New(fiber.Config{
		Views:         nil,
		Prefork:       true,
		StrictRouting: true,
		CaseSensitive: true,
		Immutable:     true,
		AppName:       "Polygon",
	})

	server.Use(cors.New(cors.ConfigDefault))
	server.Use(favicon.New(favicon.ConfigDefault))
	server.Use(compress.New(compress.ConfigDefault))

	server.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 10 * time.Minute,
	}))

	server.Use(cache.New(cache.Config{
		CacheControl:         true,
		StoreResponseHeaders: false,
		Expiration:           5 * time.Minute,
	}))

	log.Fatal(server.Listen(Configuration.Polygon.Addr))
}
