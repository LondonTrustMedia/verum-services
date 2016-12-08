// written by London Trust Media
// released under the MIT license
package ircmodes

// ModeType refers to the mode type, which change how each specific mode is set and unset.
type ModeType int

const (
	// TypeA means modes that add/remove values from a list (such as channel bans).
	TypeA ModeType = iota
	// TypeB means modes that change settings on a channel, they must always have params.
	TypeB
	// TypeC means modes that change a setting on a chan, must have param when setting but no param on unsetting.
	TypeC
	// TypeD means modes that act as a flag, without any parameter.
	TypeD
	// ChanPrefix is for a channel prefix (voice, chanop, etc).
	ChanPrefix
)

// Mode represents an IRC mode.
type Mode struct {
	Name string
	Type ModeType
}

// ModeManager stores which letters represent which modes, and which mode letters are used for them.
// It's used when parsing incoming modestrings and such.
type ModeManager struct {
	Modes      map[byte]*Mode
	NameToMode map[string]*Mode
	NameTobyte map[string]byte
}

// NewModeManager returns a fresh ModeManager.
func NewModeManager() ModeManager {
	var mm ModeManager
	mm.Modes = make(map[byte]*Mode)
	mm.NameToMode = make(map[string]*Mode)
	mm.NameTobyte = make(map[string]byte)
	return mm
}

// AddMode adds a mode to our internal list.
func (mm *ModeManager) AddMode(char byte, mode *Mode) {
	mm.Modes[char] = mode
	mm.NameToMode[mode.Name] = mode
	mm.NameTobyte[mode.Name] = char
}

// ParseModestring parses an incoming modestring and returns a ModeList based on the mode types we have.
func (mm *ModeManager) ParseModestring(params ...string) ModeList {
	ml := NewModeList()

	//TODO(dan): populate ml

	return ml
}

// ModeVal is used in storing mode values.
type ModeVal struct {
	Param string
	List  []string
}

// ModeList stores a list of modes and their values, i.e. on a client or a channel.
type ModeList struct {
	Modes map[byte]*ModeVal
}

// NewModeList returns a fresh ModeList.
func NewModeList() ModeList {
	var ml ModeList
	ml.Modes = make(map[byte]*ModeVal)
	return ml
}
