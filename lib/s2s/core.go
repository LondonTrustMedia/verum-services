// written by London Trust Media
// released under the MIT license
package s2s

// Protocol is the core S2S protocol interface that is implemented by all S2S protos.
type Protocol interface {
	Run()
	AddClient(nick, user, host, realname string)
}
