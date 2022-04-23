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
package settings

import (
	"os"

	"github.com/getpolygon/hydra"
)

const DefaultConfigFilename string = ".conf.yaml"

type LoggingEnv string

const (
	LoggingEnvProduction  LoggingEnv = "production"
	LoggingEnvDevelopment LoggingEnv = "development"
)

type SMTPSettings struct {
	User     string `yaml:"user" validate:"required" env:"POLYGON_SMTP_USER"`
	Password string `yaml:"password" validate:"required" env:"POLYGON_SMTP_PASSWORD"`
	Host     string `yaml:"host" validate:"required" validate:"uri" env:"POLYGON_SMTP_HOST"`
	From     string `yaml:"from" validate:"required" validate:"email" env:"POLYGON_SMTP_FROM"`
	Port     int16  `yaml:"port" validate:"required" validate:"number" env:"POLYGON_SMTP_PORT"`
}

type SecuritySettings struct {
	OpenRegistrations   bool   `yaml:"open_registrations" env:"POLYGON_SECURITY_OPEN_REGISTRATIONS"`
	ReportInstanceStats bool   `yaml:"report_instance_stats" env:"POLYGON_SECURITY_REPORT_INSTANCE_STATS"`
	JWTSigningKey       string `yaml:"jwt_signing_key" validate:"required" validate:"alphanum" validate:"min=64" env:"POLYGON_SECURITY_JWT_SIGNING_KEY"`
}

type Settings struct {
	SMTP     SMTPSettings     `yaml:"smtp" validate:"required"`
	Security SecuritySettings `yaml:"security" validate:"required"`
	Logging  *LoggingEnv      `yaml:"logging" env:"POLYGON_LOGGING"`
	Address  string           `yaml:"address" env:"POLYGON_BIND_ADDRESS"`
	Redis    string           `yaml:"redis" validate:"required" validate:"uri" env:"POLYGON_REDIS_URL"`
	Postgres string           `yaml:"postgres" validate:"required" validate:"uri" env:"POLYGON_POSTGRES_URL"`
}

// This function will load the configuration from the specified config file
// and environment variables, and will return a parsed settings struct.
func New() (*Settings, error) {
	var filename string
	if v := os.Getenv("POLYGON_CONFIG_FILE"); v == "" {
		// Setting the filename to a pre-defined default, if the
		// environment variable is empty.
		filename = DefaultConfigFilename
	} else {
		filename = v
	}

	// These are the default search paths where the configuration
	// of Polygon will be scanned.
	var paths []string = []string{
		".",
		"~/.getpolygon/corexp",
		"/etc/getpolygon/corexp",
	}

	if v := os.Getenv("POLYGON_CONFIG_PATH"); v != "" {
		// If a custom path IS specified via environment variables, append
		// an additional path to search index.
		paths = append(paths, v)
	}

	// Initializing a new Hydra instance, and loading the configuration
	// paths, as well as the provided filename to it.
	h := hydra.Hydra{
		Config: hydra.Config{
			Paths:    paths,
			Filename: filename,
		},
	}

	// Parsing the configuration
	s, err := h.Load(new(Settings))
	if err != nil {
		return nil, err
	}

	return s.(*Settings), nil
}
