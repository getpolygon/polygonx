package config

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Use a single instance of Validate, so that it caches
// struct info.
var validate *validator.Validate

// This is an enum, which represents the selected mailing
// interface from the configuration that is going to be
// used by default for sending emails to the users.
type EmailClient string

const (
	EmailClientNone    EmailClient = "none"
	EmailClientCourier EmailClient = "courier"
	EmailClientDefault EmailClient = "default"
)

// This is a struct, which represents the structure of the
// configuration file for Polygon. Matched files will be
// deserialized to the match the following structure.
type Config struct {
	Polygon struct {
		General *struct {
			Addr      string
			Heartbeat bool
		}

		Persistence struct {
			Redis    string `validate:"required,uri"`
			Postgres string `validate:"required,uri"`
		}

		Mail *struct {
			TemplatesLocation *string      `validate:"dir"`
			Client            *EmailClient `validate:"alpha"`
		}

		Security *struct {
			SMTP *struct {
				Secure bool `validate:"bool"`
				User   string
				Pass   string
				Port   string `validate:"number"`
				Host   string `validate:"url"`
			}

			Requests *struct {
				Max      *int `validate:"number"`
				Timespan *int `validate:"number"`
				Log      *bool
			}

			Accounts *struct {
				ForceEmailVerification *bool
			}

			JWT struct {
				// The secret must contain at least 64 characters (equivalent of 512 bits)
				Secret string `validate:"required,min=64"`
				Issuer *string
			}
		}
	}
}

func init() {
	validate = validator.New()
}

func Load() error {
	viper.SetConfigType("yaml")
	viper.SetConfigName(".conf")

	// Supported configuration directories
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/getpolygon/corexp")
	viper.AddConfigPath("$HOME/.getpolygon/corexp")

	// Setting default values for the configuration options
	// that have not been specified.
	viper.SetDefault("polygon.general.heartbeat", true)
	viper.SetDefault("polygon.general.addr", "127.0.0.1:8234")
	viper.SetDefault("polygon.security.jwt.issuer", "getpolygon")
	viper.SetDefault("polygon.security.requests.max", 200)
	viper.SetDefault("polygon.security.requests.log", false)
	viper.SetDefault("polygon.security.requests.timespan", 1*time.Minute)
	viper.SetDefault("polygon.security.accounts.forceEmailVerification", false)

	// Reading the configuration via viper
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	var config Config
	// Parsing YAML configuration file into a Go struct
	err = viper.UnmarshalExact(&config)
	if err != nil {
		return err
	}

	// Validating the configuration using the `go-playground/validator` package
	err = validate.Struct(&config)
	if err != nil {
		return err
	}

	return nil
}
