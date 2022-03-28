package config

import (
	"time"

	"github.com/spf13/viper"
)

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
		General struct {
			RequestLogging       bool
			EnableHeartbeatRoute bool
			Addr                 string
		}

		Persistence struct {
			Redis    string
			Postgres string
		}

		Security struct {
			Requests struct {
				Max      int
				Timespan int
			}

			Accounts struct {
				ForceEmailVerification bool
			}

			JWT struct {
				Secret string
				Issuer string
			}
		}
	}
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
	viper.SetDefault("polygon.general.logRequests", false)
	viper.SetDefault("polygon.general.enableHeartbeatRoute", true)
	viper.SetDefault("polygon.general.addr", "127.0.0.1:8234")
	viper.SetDefault("polygon.security.jwt.issuer", "getpolygon")
	viper.SetDefault("polygon.security.requests.max", 200)
	viper.SetDefault("polygon.security.requests.threshold", 1*time.Minute)
	viper.SetDefault("polygon.security.accounts.forceEmailVerification", false)

	// Reading the configuration
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
