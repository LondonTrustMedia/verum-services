// written by London Trust Media
// released under the MIT license
package inspircd

import (
	"time"

	"github.com/Verum/veritas/lib/s2s/deps/ircmodes"
)

// Client represents an InspIRCd client.
type Client struct {
	Server            *Server
	UID               string
	Timestamp         time.Time
	Nick              string
	Hostname          string
	DisplayedHostname string
	Ident             string
	IP                string
	SignonTime        time.Time
	Modes             map[string]string
	Realname          string
}

// MakeClient makes an InspIRCd client.
func MakeClient(uid, timestamp, nick, hostname, displayedHostname, ident, ip, signonTime string, modes ircmodes.ModeList, realname string) (*Client, error) {
	return nil, nil
	// return &c, nil
}
