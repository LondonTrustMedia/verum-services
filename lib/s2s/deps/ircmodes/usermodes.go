// written by London Trust Media
// released under the MIT license
package ircmodes

// UserMode is a mode that can be set on a user.
type UserMode Mode

var (
	// standard

	UserInvisible = UserMode{
		Name: "invisible",
		Type: TypeD,
	}
	UserOper = UserMode{
		Name: "oper",
		Type: TypeD,
	}
	UserRegistered = UserMode{
		Name: "registered",
		Type: TypeD,
	}
	UserWallops = UserMode{
		Name: "wallops",
		Type: TypeD,
	}

	// widespread

	UserBot = UserMode{
		Name: "bot",
		Type: TypeD,
	}
)
