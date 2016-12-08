// written by London Trust Media
// released under the MIT license
package lib

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config holds the core Veritas configuration information.
type Config struct {
	Server struct {
		Name        string
		Description string
		AdminEmail  string `yaml:"admin-email"`
		BannedEmail string `yaml:"ban-email"`
	}
	Linking struct {
		Module        string `yaml:"protocol-module"`
		RemoteAddress string `yaml:"remote-address"`
		UseTLS        bool   `yaml:"use-tls"`
		ServerID      string `yaml:"server-id"`
		SendPass      string `yaml:"send-password"`
		ReceivePass   string `yaml:"receive-password"`
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
