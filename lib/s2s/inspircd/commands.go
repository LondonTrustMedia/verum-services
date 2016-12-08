// written by London Trust Media
// released under the MIT license
package inspircd

import (
	"github.com/DanielOaks/girc-go/ircmsg"
)

// Command represents a command that the Insp S2S protocol module can handle.
type Command struct {
	Handler func(*InspIRCd, ircmsg.IrcMessage) error
}

var (
	// CommandHandlers are all the command handlers we implement.
	CommandHandlers = map[string]Command{}
)
