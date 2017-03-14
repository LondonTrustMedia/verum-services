// written by London Trust Media
// released under the MIT license
package inspircd

import "github.com/Verum/veritas/lib/s2s/deps/ircmodes"

// Channel represents an InspIRCd channel.
type Channel struct {
	Timestamp int64
	Name      string
	Topic     string
	Modes     ircmodes.ModeList
	Members   MemberList
}

// MemberList holds the members of the channel.
type MemberList struct {
	Members  map[*Client]bool
	Prefixes map[*Client]ircmodes.PrefixList
}

// NewMemberList returns a new MemberList.
func NewMemberList() MemberList {
	var ml MemberList
	ml.Members = make(map[*Client]bool)
	ml.Prefixes = make(map[*Client]ircmodes.PrefixList)
	return ml
}
