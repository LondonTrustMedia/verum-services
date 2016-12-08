// written by London Trust Media
// released under the MIT license
package inspircd

import (
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/DanielOaks/girc-go/ircmap"
	"github.com/DanielOaks/girc-go/ircmsg"
	"github.com/Verum/veritas/lib"
	"github.com/Verum/veritas/lib/s2s/deps"
)

// InspIRCd is the S2S protocol module for Insp.
type InspIRCd struct {
	protocol    string
	casemapping ircmap.MappingType
	s           deps.RFC1459Socket
	sid         string // server id

	receivedFirstBurst bool            // whether we've received first burst from remote
	modsupport         map[string]bool // modules loaded on remote
	chanmodes          map[string]string
	usermodes          map[string]string
	capabilities       map[string]string

	clients   map[string]*Client
	myClients map[string]*Client
}

// MakeInsp returns an InspIRCd S2S protocol module.
func MakeInsp(config *lib.Config) (*InspIRCd, error) {
	// we don't support generating SID. when we do, see how Insp does it.
	if len(config.Linking.ServerID) != 3 {
		return nil, deps.ErrorSIDIncorrect
	}

	// create protocol module
	var p InspIRCd

	p.protocol = "InspIRCd"
	p.casemapping = ircmap.RFC1459
	p.modsupport = make(map[string]bool)
	p.chanmodes = make(map[string]string)
	p.usermodes = make(map[string]string)
	p.capabilities = make(map[string]string)

	return &p, nil
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
	p.s = deps.NewRFC1459Socket(conn)
	p.s.Start()

	// insp handshake
	m, line, err := p.s.ReceiveMessage()

	if err != nil {
		return fmt.Errorf("Could not parse line [%s]: %s", line, err.Error())
	}

	if m.Command == "CAPAB" && len(m.Params) == 2 && m.Params[0] == "START" && m.Params[1] == "1202" {
		// fall-through
	} else {
		return fmt.Errorf("CAPAB START line was not correct, got [%s]", line)
	}

	// send our CAPAB burst, incoming is handled by regular command handlers
	p.s.Send(nil, "", "CAPAB", "START", "1202")
	p.s.Send(nil, "", "CAPAB", "CAPABILITIES", "PROTOCOL=1202")
	p.s.Send(nil, "", "CAPAB", "END")

	// send our SERVER line
	p.sid = config.Linking.ServerID
	p.s.Send(nil, "", "SERVER", config.Server.Name, config.Linking.SendPass, "0", p.sid, config.Server.Description)

	// send burst as well
	p.s.Send(nil, p.sid, "BURST", strconv.FormatInt(time.Now().Unix(), 10))
	p.s.Send(nil, p.sid, "VERSION", fmt.Sprintf("Veritas-%s using %s", lib.SemVer, p.protocol))
	p.s.Send(nil, p.sid, "ENDBURST")

	for {
		//TODO(dan): receive message or signal, select{} etc
		m, line, err := p.s.ReceiveMessage()
		if err == ircmsg.ErrorLineIsEmpty {
			// skip empty lines
			continue
		}
		if err == nil {
			if m.Command == "ERROR" {
				fmt.Println("Received an ERROR, disconnecting:", line)
				return fmt.Errorf("Received an ERROR from remote: %s", line)
			}

			err = HandleCommand(p, m, line)
			if err != nil {
				return fmt.Errorf("Error processing line [%s]: %s", line, err.Error())
			}
		} else {
			fmt.Println(fmt.Sprintf("Could not decode line [%s]: %s", line, err.Error()))
		}
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
