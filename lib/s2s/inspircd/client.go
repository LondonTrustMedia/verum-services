// written by London Trust Media
// released under the MIT license
package inspircd

import (
	"time"

	"strconv"

	"github.com/Verum/veritas/lib/s2s/deps/ircmodes"
)

// Client represents an InspIRCd client.
type Client struct {
	Server            *Server
	UID               string
	Timestamp         int64
	Nick              string
	Hostname          string
	DisplayedHostname string
	Ident             string
	IP                string
	SignonTime        time.Time
	Modes             ircmodes.ModeList
	Realname          string
	OperType          string

	channels map[*Channel]bool
}

// MakeClient makes an InspIRCd client.
func MakeClient(server *Server, uid, timestamp, nick, hostname, displayedHostname, ident, ip, signonTime string, modes ircmodes.ModeList, realname string) (*Client, error) {
	// make timestamp and signontime
	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return nil, err
	}

	signonInt, err := strconv.ParseInt(signonTime, 10, 64)
	if err != nil {
		return nil, err
	}
	signon := time.Unix(signonInt, 0)

	// assemble client
	c := Client{
		Server:            server,
		UID:               uid,
		Timestamp:         ts,
		Nick:              nick,
		Hostname:          hostname,
		DisplayedHostname: displayedHostname,
		Ident:             ident,
		IP:                ip,
		SignonTime:        signon,
		Modes:             modes,
		Realname:          realname,
		channels:          make(map[*Channel]bool),
	}

	return &c, nil
}
