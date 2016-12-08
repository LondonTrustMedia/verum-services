// written by London Trust Media
// released under the MIT license
package inspircd

import (
	"errors"
	"fmt"

	"strings"

	"github.com/DanielOaks/girc-go/ircmsg"
)

var (
	// ErrorNotEnoughParams is returned when there aren't enough params in the message.
	ErrorNotEnoughParams = errors.New("Not enough params")
)

// Command represents a command that the Insp S2S protocol module can handle.
type Command struct {
	MinParams int
	Handler   func(*InspIRCd, *ircmsg.IrcMessage) error
}

var (
	// Commands are all the S2S commands we implement.
	Commands = map[string]*Command{
		"BURST": &Command{
			MinParams: 1,
			Handler:   burstHandler,
		},
		"CAPAB": &Command{
			MinParams: 1,
			Handler:   capabHandler,
		},
		"PING": &Command{
			MinParams: 2,
			Handler:   pingHandler,
		},
		"SERVER": &Command{
			MinParams: 5,
			Handler:   serverHandler,
		},
		"UID": &Command{
			MinParams: 10,
			Handler:   uidHandler,
		},
		"VERSION": &Command{
			MinParams: 1,
			Handler:   versionHandler,
		},
	}
)

// HandleCommand dispatches a command, and returns an error if there is one.
func HandleCommand(p *InspIRCd, m *ircmsg.IrcMessage, line string) error {
	cmd, exists := Commands[m.Command]
	if !exists || cmd == nil {
		return fmt.Errorf("Command %s not implemented", m.Command)
	}

	// check param length, etc
	if len(m.Params) < cmd.MinParams {
		return fmt.Errorf("Not enough params for %s, expected %d but got %d: %s", m.Command, cmd.MinParams, len(m.Params), line)
	}

	return cmd.Handler(p, m)
}

// Handlers

func burstHandler(p *InspIRCd, m *ircmsg.IrcMessage) error {
	// we don't care about bursts starting, only ending
	return nil
}

func capabHandler(p *InspIRCd, m *ircmsg.IrcMessage) error {
	subcmd := m.Params[0]
	if subcmd == "MODULES" || subcmd == "END" {
		// we don't care about these
		return nil
	}

	// all from here on need 2 params
	if len(m.Params) < 2 {
		return ErrorNotEnoughParams
	}

	if subcmd == "MODSUPPORT" {
		names := strings.Split(m.Params[1], " ")
		for _, name := range names {
			if len(name) > 1 {
				p.modsupport[name] = true
			}
		}
	} else if subcmd == "CHANMODES" {
		fmt.Println("CHANMODES:", m.Params[1])
		vars := strings.Split(m.Params[1], " ")
		for _, val := range vars {
			if len(val) > 1 {
				keyval := strings.SplitN(val, "=", 2)
				key := keyval[0]
				var value string
				if len(keyval) < 2 {
					value = keyval[1]
				}

				p.chanmodes[key] = value
			}
		}
	} else if subcmd == "USERMODES" {
		vars := strings.Split(m.Params[1], " ")
		for _, val := range vars {
			if len(val) > 1 {
				keyval := strings.SplitN(val, "=", 2)
				key := keyval[0]
				var value string
				if len(keyval) < 2 {
					value = keyval[1]
				}

				p.usermodes[key] = value
			}
		}
	} else if subcmd == "CAPABILITIES" {
		vars := strings.Split(m.Params[1], " ")
		for _, val := range vars {
			if len(val) > 1 {
				keyval := strings.SplitN(val, "=", 2)
				key := keyval[0]
				var value string
				if len(keyval) < 2 {
					value = keyval[1]
				}

				p.capabilities[key] = value
			}
		}
	} else {
		fmt.Println("Unknown CAPAB:", m)
	}

	return nil
}

func pingHandler(p *InspIRCd, m *ircmsg.IrcMessage) error {
	//TODO(dan): Update our own PING time based on this?
	p.s.Send(nil, p.sid, "PING", m.Params[1], m.Params[0])
	fmt.Println("PING!!!", m)
	return nil
}

func serverHandler(p *InspIRCd, m *ircmsg.IrcMessage) error {
	newServer := Server{
		SID:         m.Params[3],
		Name:        m.Params[0],
		Description: m.Params[4],
	}

	if p.uplink == nil {
		// check password
		if m.Params[1] != p.config.Linking.ReceivePass {
			return fmt.Errorf("Receive password is not correct. Expected [%s] but got [%s]", p.config.Linking.ReceivePass, m.Params[1])
		}

		p.uplink = &newServer
	}

	p.servers[newServer.Name] = &newServer

	//TODO(dan): Dispatch NewServer event here?

	return nil
}

func uidHandler(p *InspIRCd, m *ircmsg.IrcMessage) error {
	fmt.Println("UID:", m)

	// uid := m.Params[0]
	// timestamp := m.Params[1]
	// nick := m.Params[2]
	// hostname := m.Params[3]
	// displayedHostname := m.Params[4]
	// ident := m.Params[5]
	// ip := m.Params[6]
	// signonTime := m.Params[7]
	// realname := m.Params[len(m.Params)-1]

	// // get modes
	// var modes []string
	// for i, param := range m.Params {
	// 	if i > 7 && i < len(m.Params)-1 {
	// 		modes = append(modes, param)
	// 	}
	// }

	// c := MakeClient(uid, timestamp, nick, hostname, displayedHostname, ident, ip, signonTime, modes, realname)

	// fmt.Println("Client made:", c)

	return nil
}

func versionHandler(p *InspIRCd, m *ircmsg.IrcMessage) error {
	// we don't really care about versions, at least not right now
	return nil
}
