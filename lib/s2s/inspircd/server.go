// written by London Trust Media
// released under the MIT license
package inspircd

// Server represents an InspIRCd server.
type Server struct {
	SID         string
	Name        string
	Description string

	// Links holds the server that this server is linked to.
	// Could be used to calculate stuff in netsplits/etc?
	Links map[*Server]bool
}
