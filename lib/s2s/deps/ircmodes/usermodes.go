// written by London Trust Media
// released under the MIT license
package ircmodes

// UserMode is a mode that can be set on a user.
type UserMode Mode

// UserModes
var (
	UserInvible = UserMode{
		Name: "invible",
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
)
