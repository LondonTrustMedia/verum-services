// written by London Trust Media
// released under the MIT license
package inspircd

import "github.com/Verum/veritas/lib/s2s/deps/ircmodes"

var (
	chanDelayMsg = ircmodes.Mode{
		Name: "insp-delaymsg",
		Type: ircmodes.TypeC,
	}
	chanExemptChanops = ircmodes.Mode{
		Name: "insp-exemptchanops",
		Type: ircmodes.TypeA,
	}
	chanRegInvite = ircmodes.Mode{
		Name: "insp-reginvite",
		Type: ircmodes.TypeD,
	}
	chanRegModerated = ircmodes.Mode{
		Name: "insp-regmoderated",
		Type: ircmodes.TypeD,
	}
	chanTopicLock = ircmodes.Mode{
		Name: "insp-topiclock",
		Type: ircmodes.TypeD,
	}

	// ChanModes are the channel modes that Insp uses.
	ChanModes = map[string]*ircmodes.Mode{
		"admin":         &ircmodes.ChanAdmin,
		"ban":           &ircmodes.ChanBan,
		"banexception":  &ircmodes.ChanBanException,
		"blockcolor":    &ircmodes.ChanBlockColor,
		"c_registered":  &ircmodes.ChanRegistered,
		"delaymsg":      &chanDelayMsg,
		"exemptchanops": &chanExemptChanops,
		"founder":       &ircmodes.ChanFounder,
		"halfop":        &ircmodes.ChanHalfop,
		"invex":         &ircmodes.ChanInviteException,
		"inviteonly":    &ircmodes.ChanInviteOnly,
		"joinflood":     &ircmodes.ChanJoinThrottle,
		"key":           &ircmodes.ChanKey,
		"limit":         &ircmodes.ChanLimit,
		"moderated":     &ircmodes.ChanModerated,
		"noextmsg":      &ircmodes.ChanNoExternal,
		"op":            &ircmodes.ChanOp,
		"private":       &ircmodes.ChanPrivate,
		"reginvite":     &chanRegInvite,
		"regmoderated":  &chanRegModerated,
		"secret":        &ircmodes.ChanSecret,
		"sslonly":       &ircmodes.ChanTLSOnly,
		"topiclock":     &chanTopicLock,
		"voice":         &ircmodes.ChanVoice,
	}
)

var (
	userHelpOp = ircmodes.Mode{
		Name: "insp-helpop",
		Type: ircmodes.TypeD,
	}
	userHideChans = ircmodes.Mode{
		Name: "insp-hidechans",
		Type: ircmodes.TypeD,
	}
	userRegDeaf = ircmodes.Mode{
		Name: "insp-regdeaf",
		Type: ircmodes.TypeD,
	}
	userSnoMask = ircmodes.Mode{
		Name: "insp-snomask",
		Type: ircmodes.TypeA,
	}

	// UserModes are the user modes that Insp uses.
	UserModes = map[string]*ircmodes.Mode{
		"bot":          &ircmodes.UserBot,
		"helpop":       &userHelpOp,
		"hidechans":    &userHideChans,
		"invisible":    &ircmodes.UserInvisible,
		"oper":         &ircmodes.UserOper,
		"regdeaf":      &userRegDeaf,
		"snomask":      &userSnoMask,
		"u_registered": &ircmodes.UserRegistered,
		"wallops":      &ircmodes.UserWallops,
	}
)
