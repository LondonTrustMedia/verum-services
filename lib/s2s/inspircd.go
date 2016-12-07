// written by London Trust Media
// released under the MIT license
package s2s

import (
	"crypto/tls"
	"fmt"
	"net"

	"github.com/DanielOaks/girc-go/ircmap"
	"github.com/Verum/veritas/lib"
)

// InspIRCd is the S2S protocol module for Insp.
type InspIRCd struct {
	protocol    string
	casemapping ircmap.MappingType
	s           RFC1459Socket
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

// Run connects to the remote and starts running.
func (p *InspIRCd) Run(config *lib.Config) error {
	// connect
	var conn net.Conn
	var err error

	if config.Linking.UseTLS {
		conn, err = tls.Dial("tcp", config.Linking.RemoteAddress, nil)
	} else {
		conn, err = net.Dial("tcp", config.Linking.RemoteAddress)
	}

	if err != nil {
		return err
	}

	// open socket properly
	p.s = NewRFC1459Socket(conn)
	p.s.Start()

	for {
		fmt.Println("LINE:", <-p.s.ReceiveLines)
	}

	return nil
}

// CasemapString returns a casemapped version of the source string.
func (p *InspIRCd) CasemapString(source string) (string, error) {
	return ircmap.Casefold(p.casemapping, source)
}

func (p *InspIRCd) AddClient(nick, user, host, realname string) error {
	//TODO(dan): run here
	return nil
}
