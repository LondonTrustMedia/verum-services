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

	Handler func(*InspIRCd, *ircmsg.IrcMessage) error
}

var (
	// Commands are all the S2S commands we implement.
	Commands = map[string]*Command{
		"CAPAB": &Command{
			MinParams: 1,
			Handler:   capabHandler,
		},
		"PING": &Command{
			MinParams: 2,
			Handler:   pingHandler,
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
