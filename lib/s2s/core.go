// written by London Trust Media
// released under the MIT license
package s2s

import (
	"errors"
	"strings"

	"github.com/Verum/veritas/lib"
)

var (
	ErrorNoProtocol = errors.New("Protocol not found")
)

// Protocol is the core S2S protocol interface that is implemented by all S2S protos.
type Protocol interface {
	Run()
	AddClient(nick, user, host, realname string)
}

// MakeProto returns a protocol module given the config.
func MakeProto(config *lib.Config) (*Protocol, error) {
	var p *Protocol

	protoName := strings.ToLower(config.IRCd.Module)
	if protoName == "inspircd" {
		inspProto, err := makeInsp(config)
		return inspProto.(*Protocol), err
	}

	return nil, ErrorNoProtocol
}
