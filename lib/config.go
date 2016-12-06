// written by London Trust Media
// released under the MIT license
package lib

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config holds the core Veritas configuration information.
type Config struct {
	IRCd struct {
		Name     string
		Module   string `yaml:"protocol-module"`
		ServerID string `yaml:"sid"`
	}
}

// LoadConfig returns a Config object or returns an error.
func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config *Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	// confirm details
	return config, nil
}
