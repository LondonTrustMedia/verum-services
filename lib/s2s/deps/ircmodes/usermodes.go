// written by London Trust Media
// released under the MIT license
package ircmodes

var (
	// standard

	UserInvisible = Mode{
		Name: "invisible",
		Type: TypeD,
	}
	UserOper = Mode{
		Name: "oper",
		Type: TypeD,
	}
	UserRegistered = Mode{
		Name: "registered",
		Type: TypeD,
	}
	UserWallops = Mode{
		Name: "wallops",
		Type: TypeD,
	}

	// widespread

	UserBot = Mode{
		Name: "bot",
		Type: TypeD,
	}
)
