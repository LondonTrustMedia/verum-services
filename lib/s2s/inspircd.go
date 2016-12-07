// written by London Trust Media
// released under the MIT license
package s2s

import "github.com/DanielOaks/girc-go/ircmap"
import "github.com/Verum/veritas/lib"

// InspIRCd is the S2S protocol module for Insp.
type InspIRCd struct {
	Protocol    string
	Casemapping ircmap.MappingType
}

func makeInsp(config *lib.Config) (InspIRCd, error) {
	var p InspIRCd

	p.Protocol = "InspIRCd"
	p.Casemapping = ircmap.RFC1459

	return p, nil
}

func (p *InspIRCd) Run() {
	//TODO(dan): run here
}

func (p *InspIRCd) AddClient(nick, user, host, realname string) error {
	//TODO(dan): run here
	return nil
}
