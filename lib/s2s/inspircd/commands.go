// written by London Trust Media
// released under the MIT license
package inspircd

import (
	"errors"
	"fmt"
	"strconv"

	"strings"

	"github.com/DanielOaks/girc-go/ircmsg"
	"github.com/Verum/veritas/lib/s2s/deps/ircmodes"
	"github.com/davecgh/go-spew/spew"
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
		"FJOIN": &Command{
			MinParams: 4,
			Handler:   fjoinHandler,
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
		fmt.Println(fmt.Sprintf("Command %s not implemented", m.Command))
		return nil
	}

	// check param length, etc
	if len(m.Params) < cmd.MinParams {
		fmt.Println(fmt.Sprintf("Not enough params for %s, expected %d but got %d: %s", m.Command, cmd.MinParams, len(m.Params), line))
		return nil
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

				if len(value) == 1 {
					// regular modes are only 1 char long
					mode, exists := ChanModes[key]
					if exists {
						p.chanmodes.AddMode(value[0], mode)
					} else {
						fmt.Println("I don't know mode", val, "autogenerating")
						p.unknownCmodes[value[0]] = key
					}
				} else if len(value) == 2 {
					// channel prefixes
					mode, exists := ChanModes[key]
					if exists {
						p.chanmodes.AddMode(value[1], mode)
					} else {
						return fmt.Errorf("I can't parse mode %s correctly: %s", key, m.Params[1])
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

				if key == "PREFIX" {
					fmt.Println("TODO(dan): Track channel prefixes properly.")
				}

				p.capabilities[key] = value
			}
		}
	} else {
		fmt.Println("Unknown CAPAB:", m)
	}

	return nil
}

func fjoinHandler(p *InspIRCd, m *ircmsg.IrcMessage) error {
	fmt.Println("FJOIN:", m)

	// get name and timestamp
	casefoldedName, err := p.CasemapString(m.Params[0])
	if err != nil {
		fmt.Println("Could not casemap given channel name:", m)
		return ErrorUnknown
	}

	ts, err := strconv.ParseInt(m.Params[1], 10, 64)
	if err != nil {
		return err
	}

	// get modes
	var modes []string
	for i, param := range m.Params {
		if i > 1 && i < len(m.Params)-1 {
			modes = append(modes, param)
		}
	}

	modeList := p.chanmodes.ParseModestringToList(modes...)
	members := NewMemberList()

	// add chan priv info
	// last param contains user info
	for _, upriv := range strings.Split(m.Params[len(m.Params)-1], " ") {
		if len(upriv) < 1 {
			continue
		}

		info := strings.Split(upriv, ",")
		uid := info[1]
		c := p.clients[uid]

		// add to members list
		members.Members[c] = true

		// add to prefixes list
		prefixes, exists := members.Prefixes[c]
		if !exists {
			prefixes = make(ircmodes.PrefixList)
		}
		for _, char := range info[0] {
			mode := p.chanmodes.Modes[byte(char)]
			prefixes[mode] = true
		}
		members.Prefixes[c] = prefixes
	}

	// add to internal list
	ch, exists := p.channels[casefoldedName]
	if exists {
		// update channel info with given info, and modes if ts is newer
	} else {
		ch = &Channel{
			Name:      m.Params[0],
			Timestamp: ts,
			Modes:     modeList,
			Members:   members,
		}
	}

	spew.Dump(ch)

	p.channels[casefoldedName] = ch

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
	p.s.Send(nil, p.sid, "PONG", m.Params[1], m.Params[0])
	return nil
}

func serverHandler(p *InspIRCd, m *ircmsg.IrcMessage) error {
	newServer := Server{
		SID:         m.Params[3],
		Name:        m.Params[0],
		Description: m.Params[4],
		Links:       make(map[*Server]bool),
	}

	if p.uplink == nil {
		// check password
		if m.Params[1] != p.config.Linking.ReceivePass {
			return fmt.Errorf("Receive password is not correct. Expected [%s] but got [%s]", p.config.Linking.ReceivePass, m.Params[1])
		}

		p.uplink = &newServer

		p.thisServer.Links[&newServer] = true
	} else {
		fmt.Println("I don't know how to handle incoming links from the remote side:", m)
		fmt.Println("Please make sure that I'm added to the Links block of the relevant server.")
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
