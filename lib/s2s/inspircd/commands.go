// written by London Trust Media
// released under the MIT license
package inspircd

import (
	"errors"
	"fmt"

	"strings"

	"github.com/DanielOaks/girc-go/ircmsg"
	"github.com/Verum/veritas/lib/s2s/deps/ircmodes"
)

var (
	// ErrorNotEnoughParams is returned when there aren't enough params in the message.
	ErrorNotEnoughParams = errors.New("Not enough params")
	// ErrorUnknown represents an unknown error.
	ErrorUnknown = errors.New("Unknown error, check logs")
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
		"OPERTYPE": &Command{
			MinParams: 1,
			Handler:   opertypeHandler,
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
		//TODO(dan): store channel prefixes as well.
		vars := strings.Split(m.Params[1], " ")
		for _, val := range vars {
			if len(val) > 1 {
				keyval := strings.SplitN(val, "=", 2)
				key := keyval[0]
				var value string
				if len(keyval) > 1 {
					value = keyval[1]
				}

				// regular modes are only 1 char long
				if len(value) == 1 {
					mode, exists := ChanModes[key]
					if exists {
						p.chanmodes.AddMode(value[0], mode)
					} else {
						fmt.Println("I don't know mode", val, "autogenerating")
						p.unknownCmodes[value[0]] = key
					}
				}

				p.chanmodesraw[key] = value
			}
		}
	} else if subcmd == "USERMODES" {
		vars := strings.Split(m.Params[1], " ")
		for _, val := range vars {
			if len(val) > 1 {
				keyval := strings.SplitN(val, "=", 2)
				key := keyval[0]
				var value string
				if len(keyval) > 1 {
					value = keyval[1]
				}

				// regular modes are only 1 char long
				if len(value) == 1 {
					mode, exists := UserModes[key]
					if exists {
						p.usermodes.AddMode(value[0], mode)
					} else {
						fmt.Println("I don't know mode", val, "autogenerating")
						p.unknownUmodes[value[0]] = key
					}
				}

				p.usermodesraw[key] = value
			}
		}
	} else if subcmd == "CAPABILITIES" {
		fmt.Println("CAPABILITIES:", m.Params[1])
		vars := strings.Split(m.Params[1], " ")
		for _, val := range vars {
			if len(val) > 1 {
				keyval := strings.SplitN(val, "=", 2)
				key := keyval[0]
				var value string
				if len(keyval) > 1 {
					value = keyval[1]
				}

				// autogenerate unknown channel modes
				if key == "CHANMODES" {
					splits := strings.Split(value, ",")

					var modeType ircmodes.ModeType
					for i, chars := range splits {
						if i == 0 {
							modeType = ircmodes.TypeA
						} else if i == 1 {
							modeType = ircmodes.TypeB
						} else if i == 2 {
							modeType = ircmodes.TypeC
						} else if i == 3 {
							modeType = ircmodes.TypeD
						} else {
							break
						}

						for _, char := range chars {
							name, exists := p.unknownCmodes[byte(char)]
							if exists {
								newMode := ircmodes.Mode{
									Name: fmt.Sprintf("insp-%s", name),
									Type: modeType,
								}
								p.chanmodes.AddMode(byte(char), &newMode)
							}
						}
					}
				}

				// autogenerate unknown user modes
				if key == "USERMODES" {
					splits := strings.Split(value, ",")

					var modeType ircmodes.ModeType
					for i, chars := range splits {
						if i == 0 {
							modeType = ircmodes.TypeA
						} else if i == 1 {
							modeType = ircmodes.TypeB
						} else if i == 2 {
							modeType = ircmodes.TypeC
						} else if i == 3 {
							modeType = ircmodes.TypeD
						} else {
							break
						}

						for _, char := range chars {
							name, exists := p.unknownUmodes[byte(char)]
							if exists {
								newMode := ircmodes.Mode{
									Name: fmt.Sprintf("insp-%s", name),
									Type: modeType,
								}
								p.usermodes.AddMode(byte(char), &newMode)
							}
						}
					}
				}

				p.capabilities[key] = value
			}
		}
	} else {
		fmt.Println("Unknown CAPAB:", m)
	}

	return nil
}

func opertypeHandler(p *InspIRCd, m *ircmsg.IrcMessage) error {
	c, exists := p.clients[m.Prefix]
	if !exists {
		fmt.Println("Got OPERTYPE", m.Params, "for unknown client", m.Prefix)
		return ErrorUnknown
	}

	c.OperType = m.Params[0]
	p.clients[m.Prefix] = c
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
	// get server
	server := p.servers[m.Prefix]

	uid := m.Params[0]
	timestamp := m.Params[1]
	nick := m.Params[2]
	hostname := m.Params[3]
	displayedHostname := m.Params[4]
	ident := m.Params[5]
	ip := m.Params[6]
	signonTime := m.Params[7]
	realname := m.Params[len(m.Params)-1]

	// get modes
	var modes []string
	for i, param := range m.Params {
		if i > 7 && i < len(m.Params)-1 {
			modes = append(modes, param)
		}
	}

	modeList := p.usermodes.ParseModestringToList(modes...)

	c, err := MakeClient(server, uid, timestamp, nick, hostname, displayedHostname, ident, ip, signonTime, modeList, realname)
	if err != nil {
		return err
	}

	p.clients[uid] = c
	return nil
}

func versionHandler(p *InspIRCd, m *ircmsg.IrcMessage) error {
	// we don't really care about versions, at least not right now
	return nil
}
