// written by London Trust Media
// released under the MIT license
package s2s

import (
	"strings"

	"github.com/Verum/veritas/lib"
)

// Protocol is the core S2S protocol interface that is implemented by all S2S protos.
type Protocol interface {
	Run()
	AddClient(nick, user, host, realname string)
}

// MakeProro returns a protocol module given the config.
func MakeProro(config *lib.Config) (Protocol, error) {
	var p *Protocol

	protoName := strings.ToLower(config.IRCd.Module)
	if protoName == "inspircd" {
		inspProto, err := makeInsp(config)
		return inspProto, err
	}

	return p, nil
}
