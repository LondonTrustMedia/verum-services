package s2s

type Protocol interface {
	Run()
	AddClient(nick, user, host, realname string)
}
