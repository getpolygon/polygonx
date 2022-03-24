package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
	"polygon.am/core/pkg/types"
)

// This function will attempt to read the configuration from the supplied
// path and return the parsed struct.
func ParseConfig(path string) (*types.Config, error) {
	config, err := ioutil.ReadFile(path)
	if err == nil {
		// Creating a configuration struct with overloaded fields
		// to assign by default when the option is not provided from
		// the configuration file.
		var output types.Config = types.Config{
			Polygon: types.Polygon{
				Addr: "127.0.0.1:5000",
			},
		}

		err := yaml.Unmarshal(config, &output)
		if err == nil {
			return &output, nil
		}

		return nil, err

	}

	return nil, err
}
