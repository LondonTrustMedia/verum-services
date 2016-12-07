// written by London Trust Media
// released under the MIT license
package s2s

import "github.com/DanielOaks/girc-go/ircmap"
import "github.com/Verum/veritas/lib"

// InspIRCd is the S2S protocol module for Insp.
type InspIRCd struct {
	Casemapping ircmap.MappingType
}

func makeInsp(config *lib.Config) (*InspIRCd, error) {

}

func (p *InspIRCd) Run() {
	//TODO(dan): run here
}

func (p *InspIRCd) AddClient(nick, user, host, realname string) error {
	//TODO(dan): run here
	return nil
}
