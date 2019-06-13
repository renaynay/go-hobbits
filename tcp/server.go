// Package tcp is a framework for a TCP server that will be able to handle TCP transport
// It will also decode inbound messages and encode outbound messages
package tcp

import (
	"errors"
	"fmt"
	"net"

	"github.com/renaynay/go-hobbits/encoding"
)

type callback func(net.Conn, encoding.Message)

type Server struct {
	host string
	port int
}

// NewServer creates a new server
func NewServer(host string, port int) *Server {
	return &Server{host: host, port: port}
}

// Listen listens for incoming connections.
func (s *Server) Listen(c callback) error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		return errors.New(fmt.Sprintf("Error listening: %s", (err.Error())))
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}

		go s.handle(conn, c)
	}
}

// handle handles incoming requests.
func (*Server) handle(conn net.Conn, c callback) error {
	buf := make([]byte, 1024) // TODO: do this better

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
func (*Server) SendMessage(conn net.Conn, message encoding.Message) error {
	encoded, err := encoding.Marshal(message)
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte(encoded))
	if err != nil {
		return err
	}

	err = conn.Close()
	if err != nil {
		return err
	}

	return nil
}
