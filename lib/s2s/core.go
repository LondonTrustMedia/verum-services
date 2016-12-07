// written by London Trust Media
// released under the MIT license
package s2s

import (
	"errors"
	"strings"

	"github.com/Verum/veritas/lib"
)

var (
	// ErrorNoProtocol is what it says on the tin.
	ErrorNoProtocol = errors.New("Protocol not found")
)

// Protocol is the core S2S protocol interface that is implemented by all S2S protos.
type Protocol interface {
	// info methods
	GetProtocolName() string

	// events
	Run()

	// protocol handling/management
	CasemapString(source string) (string, error)
	AddClient(nick, user, host, realname string) error
}

// MakeProto returns a generic protocol module given the config.
func MakeProto(config *lib.Config) (Protocol, error) {
	protoName := strings.ToLower(config.IRCd.Module)

	if protoName == "inspircd" {
		inspProto, err := MakeInsp(config)
		return &inspProto, err
	}

	return nil, ErrorNoProtocol
}
