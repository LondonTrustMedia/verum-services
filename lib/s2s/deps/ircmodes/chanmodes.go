// written by London Trust Media
// released under the MIT license
package ircmodes

// ChannelMode is a mode that can be set on a channel.
type ChannelMode Mode

var (
	// prefix modes

	ChanFounder = ChannelMode{
		Name: "founder",
		Type: ChanPrefix,
	}
	ChanAdmin = ChannelMode{
		Name: "admin",
		Type: ChanPrefix,
	}
	ChanOp = ChannelMode{
		Name: "op",
		Type: ChanPrefix,
	}
	ChanHalfop = ChannelMode{
		Name: "halfop",
		Type: ChanPrefix,
	}
	ChanVoice = ChannelMode{
		Name: "voice",
		Type: ChanPrefix,
	}

	// standard

	ChanBanException = ChannelMode{
		Name: "ban-exception",
		Type: TypeA,
	}
	ChanBan = ChannelMode{
		Name: "ban",
		Type: TypeA,
	}
	ChanInviteException = ChannelMode{
		Name: "invite-exception",
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
		Name: "noextmsg",
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

	ChanJoinThrottle = ChannelMode{
		Name: "join-throttle",
		Type: TypeC,
	}
	ChanRegistered = ChannelMode{
		Name: "registered",
		Type: TypeD,
	}
	ChanRegOnly = ChannelMode{
		Name: "registered-only",
		Type: TypeD,
	}
	ChanTLSOnly = ChannelMode{
		Name: "tls-only",
		Type: TypeD,
	}
)
