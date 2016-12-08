// written by London Trust Media
// uses code from Jeremy Latt
// released under the MIT license
package s2s

import (
	"bufio"
	"errors"
	"io"
	"net"
	"strings"

	"github.com/DanielOaks/girc-go/ircmsg"
	"github.com/Verum/veritas/lib"
)

var (
	// ErrorNoProtocol is what it says on the tin.
	ErrorNoProtocol = errors.New("Protocol not found")
	// ErrorSIDIncorrect means that the SID wasn't defined or was incorrect.
	ErrorSIDIncorrect = errors.New("ServerID is either incorrect or not defined")
)

// Protocol is the core S2S protocol interface that is implemented by all S2S protos.
type Protocol interface {
	// info methods
	GetProtocolName() string

	// events
	Run(config *lib.Config) error

	// protocol handling/management
	CasemapString(source string) (string, error)
	AddClient(nick, user, host, realname string) error
}

// MakeProto returns a generic protocol module given the config.
func MakeProto(config *lib.Config) (Protocol, error) {
	protoName := strings.ToLower(config.Linking.Module)

	if protoName == "inspircd" {
		inspProto, err := MakeInsp(config)
		return inspProto, err
	}

	return nil, ErrorNoProtocol
}

// Socket represents an IRC socket.
type Socket struct {
	Closed bool
	conn   net.Conn
	reader *bufio.Reader
}

// newSocket returns a new Socket.
func newSocket(conn net.Conn) Socket {
	return Socket{
		conn:   conn,
		reader: bufio.NewReader(conn),
	}
}

// Close stops a Socket from being able to send/receive any more data.
func (socket *Socket) Close() {
	if socket.Closed {
		return
	}
	socket.Closed = true
	socket.conn.Close()
}

// Read returns a single IRC line from a Socket.
func (socket *Socket) Read() (string, error) {
	if socket.Closed {
		return "", io.EOF
	}

	lineBytes, err := socket.reader.ReadBytes('\n')

	// convert bytes to string
	line := string(lineBytes[:])

	// read last message properly, just fail next reads/writes
	if err == io.EOF {
		socket.Close()
	}

	if err == io.EOF && strings.TrimSpace(line) != "" {
		// don't do anything
	} else if err != nil {
		return "", err
	}

	return strings.TrimRight(line, "\r\n"), nil
}

// Write sends the given string out of Socket.
func (socket *Socket) Write(data string) error {
	if socket.Closed {
		return io.EOF
	}

	// write data
	_, err := socket.conn.Write([]byte(data))
	if err != nil {
		socket.Close()
		return err
	}
	return nil
}

// WriteLine writes the given line out of Socket.
func (socket *Socket) WriteLine(line string) error {
	return socket.Write(line + "\r\n")
}

// RFC1459Socket listens to a socket using the IRC protocol, processes events,
// and also sends IRC lines out of that socket. Lots of S2S protos are based
// off the RFC1459 S2S proto, so this is useful.
type RFC1459Socket struct {
	ReceiveLines chan string
	SendLines    chan string
	socket       Socket
}

// NewRFC1459Socket returns a new RFC1459Socket.
func NewRFC1459Socket(conn net.Conn) RFC1459Socket {
	return RFC1459Socket{
		ReceiveLines: make(chan string),
		SendLines:    make(chan string),
		socket:       newSocket(conn),
	}
}

// Start creates and starts running the necessary event loops.
func (rs *RFC1459Socket) Start() {
	go rs.RunSocketSender()
	go rs.RunSocketListener()
}

// RunSocketSender sends lines to the IRC socket.
func (rs *RFC1459Socket) RunSocketSender() {
	var err error
	var line string
	for {
		line = <-rs.SendLines
		err = rs.socket.Write(line)
		if err != nil {
			break
		}
	}
}

// RunSocketListener receives lines from the IRC socket.
func (rs *RFC1459Socket) RunSocketListener() {
	var errConn error
	var line string

	for {
		line, errConn = rs.socket.Read()
		rs.ReceiveLines <- line
		if errConn != nil {
			break
		}
	}
	if !rs.socket.Closed {
		rs.socket.Close()
	}
}

// ReceiveMessage receives an IRC line from remote, decodes it as an IRC message, and returns the message, the raw line, and any possible errors.
func (rs *RFC1459Socket) ReceiveMessage() (*ircmsg.IrcMessage, string, error) {
	line := <-rs.ReceiveLines
	m, err := ircmsg.ParseLine(line)
	if err != nil {
		return nil, line, err
	}

	return &m, line, nil
}

// Send sends an IRC line to the listener.
func (rs *RFC1459Socket) Send(tags *map[string]ircmsg.TagValue, prefix string, command string, params ...string) error {
	ircmsg := ircmsg.MakeMessage(tags, prefix, command, params...)
	line, err := ircmsg.Line()
	if err != nil {
		return err
	}
	rs.SendLines <- line
	return nil
}
