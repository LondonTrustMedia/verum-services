// written by London Trust Media
// released under the MIT license
package ircmodes

// PrefixList represents channel priv prefixes.
type PrefixList map[*Mode]bool

// Has returns whether or not this PrefixList has the given priv.
func (pl PrefixList) Has(priv *Mode) bool {
	return pl[priv]
}

// HasAtLeast returns whether this PrefixList contains this mode or one of a higher priv.
func (pl PrefixList) HasAtLeast(priv *Mode) bool {
	var enabled bool

	for _, prefix := range ChanPrefixes {
		if priv == prefix {
			enabled = true
		}
		if enabled && pl[prefix] {
			return true
		}
	}

	return false
}

// IngestModeChanges ingests mode changes and changes our state to accommodate.
func (pl *PrefixList) IngestModeChanges(mc *ModeChanges) {

}
