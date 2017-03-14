// written by London Trust Media
// released under the MIT license
package ircmodes

var (
	// prefix modes

	ChanFounder = Mode{
		Name: "founder",
		Type: ChanPrefix,
	}
	ChanAdmin = Mode{
		Name: "admin",
		Type: ChanPrefix,
	}
	ChanOp = Mode{
		Name: "op",
		Type: ChanPrefix,
	}
	ChanHalfop = Mode{
		Name: "halfop",
		Type: ChanPrefix,
	}
	ChanVoice = Mode{
		Name: "voice",
		Type: ChanPrefix,
	}
	ChanPrefixes = []*Mode{&ChanVoice, &ChanHalfop, &ChanOp, &ChanAdmin, &ChanFounder}

	// standard

	ChanBanException = Mode{
		Name: "ban-exception",
		Type: TypeA,
	}
	ChanBan = Mode{
		Name: "ban",
		Type: TypeA,
	}
	ChanInviteException = Mode{
		Name: "invite-exception",
		Type: TypeA,
	}
	ChanInviteOnly = Mode{
		Name: "invite-only",
		Type: TypeD,
	}
	ChanKey = Mode{
		Name: "key",
		Type: TypeB,
	}
	ChanLimit = Mode{
		Name: "limit",
		Type: TypeC,
	}
	ChanModerated = Mode{
		Name: "moderated",
		Type: TypeD,
	}
	ChanNoExternal = Mode{
		Name: "noextmsg",
		Type: TypeD,
	}
	ChanProtectedTopic = Mode{
		Name: "protected-topic",
		Type: TypeD,
	}
	ChanSecret = Mode{
		Name: "secret",
		Type: TypeD,
	}

	// widespread

	ChanJoinThrottle = Mode{
		Name: "join-throttle",
		Type: TypeC,
	}
	ChanPrivate = Mode{
		Name: "private",
		Type: TypeD,
	}
	ChanRegistered = Mode{
		Name: "registered",
		Type: TypeD,
	}
	ChanRegOnly = Mode{
		Name: "registered-only",
		Type: TypeD,
	}
	ChanTLSOnly = Mode{
		Name: "tls-only",
		Type: TypeD,
	}

	// others that are especially useful to keep synced between IRCds

	ChanBlockColor = Mode{
		Name: "block-color",
		Type: TypeC,
	}
)
