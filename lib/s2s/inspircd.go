// written by London Trust Media
// released under the MIT license
package s2s

import "github.com/DanielOaks/girc-go/ircmap"
import "github.com/Verum/veritas/lib"

// InspIRCd is the S2S protocol module for Insp.
type InspIRCd struct {
	protocol    string
	casemapping ircmap.MappingType
}

// MakeInsp returns an InspIRCd S2S protocol module.
func MakeInsp(config *lib.Config) (InspIRCd, error) {
	var p InspIRCd

	p.protocol = "InspIRCd"
	p.casemapping = ircmap.RFC1459

	return p, nil
}

// GetProtocolName returns the name of this protocol module.
func (p *InspIRCd) GetProtocolName() string {
	return p.protocol
}

func (p *InspIRCd) Run() {
	//TODO(dan): run here
}

// CasemapString returns a casemapped version of the source string.
func (p *InspIRCd) CasemapString(source string) (string, error) {
	return ircmap.Casefold(p.casemapping, source)
}

func (p *InspIRCd) AddClient(nick, user, host, realname string) error {
	//TODO(dan): run here
	return nil
}
