// Package tcp is a framework for a TCP server that will be able to handle TCP transport
// It will also decode inbound messages and encode outbound messages
package tcp

import (
	"errors"
	"fmt"
	"net"

	"github.com/renaynay/go-hobbits/encoding"
)

// Callback is a function for message handling
type Callback func(net.Conn, encoding.Message)

type Server struct {
	host string
	port int
}

// NewServer creates a new server
func NewServer(host string, port int) *Server {
	return &Server{host: host, port: port}
}

// Listen listens for incoming connections
func (s *Server) Listen(c Callback) error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		return errors.New(fmt.Sprintf("Error listening: %s", (err.Error()))) // TODO: errorf, use errors.New when it's a constant (make it public and exported
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept() // TODO: maybe the error is worth a log or an error chan since it stops the routine if you get one error
		if err != nil {
			return err
		}

		go handle(conn, c)
	}
}

// handle handles incoming requests
func handle(conn net.Conn, c Callback) error {
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		return errors.New(fmt.Sprintf("Error reading: %s", err.Error())) // TODO: clean up error
	}

	decoded, err := encoding.Unmarshal(string(buf))
	if err != nil {
		return err
	}

	go c(conn, *decoded)

	return nil
}

// SendMessage sends an encoded message
func (*Server) SendMessage(conn net.Conn, message encoding.Message) error { //TODO: how can this be easier to use? does it need to operate on Server
	defer conn.Close()

	encoded, err := encoding.Marshal(message)
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte(encoded))
	if err != nil {
		return err
	}

	return nil
}
