// written by London Trust Media
// released under the MIT license
package ircmodes

// ChannelMode is a mode that can be set on a channel.
type ChannelMode Mode

// ChanModes
var (
	// standard

	ChanBanExempt = ChannelMode{
		Name: "ban-exempt",
		Type: TypeA,
	}
	ChanBans = ChannelMode{
		Name: "ban",
		Type: TypeA,
	}
	ChanInviteExempt = ChannelMode{
		Name: "invite-exempt",
		Type: TypeA,
	}
	ChanInviteOnly = ChannelMode{
		Name: "invite-only",
		Type: TypeD,
	}
	ChanKey = ChannelMode{
		Name: "key",
		Type: TypeB,
	}
	ChanLimit = ChannelMode{
		Name: "limit",
		Type: TypeC,
	}
	ChanModerated = ChannelMode{
		Name: "moderated",
		Type: TypeD,
	}
	ChanNoExternal = ChannelMode{
		Name: "no-external",
		Type: TypeD,
	}
	ChanProtectedTopic = ChannelMode{
		Name: "protected-topic",
		Type: TypeD,
	}
	ChanSecret = ChannelMode{
		Name: "secret",
		Type: TypeD,
	}

	// widespread

	ChanTLSOnly = ChannelMode{
		Name: "tls-only",
		Type: TypeD,
	}
)
