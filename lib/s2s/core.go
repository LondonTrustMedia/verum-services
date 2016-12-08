// written by London Trust Media
// released under the MIT license
package s2s

import (
	"strings"

	"github.com/Verum/veritas/lib"
	"github.com/Verum/veritas/lib/s2s/deps"
	"github.com/Verum/veritas/lib/s2s/inspircd"
)

// Protocol is the core S2S protocol interface that is implemented by all S2S protos.
type Protocol interface {
	// info methods
	GetProtocolName() string

	// events
	Run(config *lib.Config) error

	// protocol handling/management
	CasemapString(source string) (string, error)
	AddClient(nick, user, host, realname string) error
}

// MakeProto returns a generic protocol module given the config.
func MakeProto(config *lib.Config) (Protocol, error) {
	protoName := strings.ToLower(config.Linking.Module)

	if protoName == "inspircd" {
		inspProto, err := inspircd.MakeInsp(config)
		return inspProto, err
	}

	return nil, deps.ErrorNoProtocol
}
