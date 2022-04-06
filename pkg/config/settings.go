// BSD 3-Clause License

// Copyright (c) 2021, Michael Grigoryan
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:

// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.

// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.

// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package config

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Use a single instance of Validate, so that it caches
// struct info.
var validate = validator.New()

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
			TemplatesLocation *string `validate:"dir"`
			Client            *EmailClient
		}

		Security *struct {
			SMTP *struct {
				User   string
				Pass   string
				Secure bool   `validate:"boolean"`
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

// This function will attempt to load the YAML configuration
// from the supported directories. If something fails or is
// invalid, the function will throw an error.
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
	viper.SetDefault("polygon.security.requests.max", 200)
	viper.SetDefault("polygon.security.requests.log", false)
	viper.SetDefault("polygon.security.jwt.issuer", "getpolygon")
	viper.SetDefault("polygon.security.requests.timespan", 1*time.Minute)
	viper.SetDefault("polygon.security.accounts.forceEmailVerification", false)

	// Reading the configuration via viper
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	var config Config
	// Parsing YAML configuration file into a Go struct
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	// Validating the configuration using the `go-playground/validator` package
	if err := validate.Struct(&config); err != nil {
		return err
	}

	return nil
}
