// written by London Trust Media
// released under the MIT license
package s2s

type Protocol interface {
	Run()
	AddClient(nick, user, host, realname string)
}
