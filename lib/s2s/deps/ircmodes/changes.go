// written by London Trust Media
// released under the MIT license
package ircmodes

import "fmt"

// ModeChangeOperation represents a mode change operation
type ModeChangeOperation int

const (
	Add ModeChangeOperation = iota
	Remove
)

// ModeChange represents a single mode change.
type ModeChange struct {
	Op    ModeChangeOperation
	Mode  byte
	Param string
}

// ModeChanges represents a series of mode changes.
type ModeChanges []ModeChange

// ParseModeChanges parses an incoming modestring and returns a ModeChanges based on the mode types we have.
func (mm *ModeManager) ParseModeChanges(params ...string) ModeChanges {
	var mc ModeChanges

	op := Add
	paramPointer := 1

	for _, char := range params[0] {
		if char == '+' {
			op = Add
			continue
		} else if char == '-' {
			op = Remove
			continue
		}

		mode := mm.Modes[byte(char)]
		if mode == nil {
			// unknown mode, this should not happen
			fmt.Println("Unknown mode char found:", char, "in", params)
			continue
		}

		change := ModeChange{
			Op:   op,
			Mode: byte(char),
		}

		if mode.Type == TypeA || mode.Type == TypeB || (mode.Type == TypeC && op == Add) {
			if paramPointer < len(params) {
				change.Param = params[paramPointer]
				paramPointer++
			} else {
				// no param for this mode that requires one, so continue
				continue
			}
		}

		// and add the mode change
		mc = append(mc, change)
	}

	return mc
}
