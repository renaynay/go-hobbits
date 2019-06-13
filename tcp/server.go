// Package tcp is a framework for a TCP server that will be able to handle TCP transport
// It will also decode inbound messages and encode outbound messages
package tcp

import (
	"errors"
	"fmt"
	"net"

	"github.com/renaynay/go-hobbits/encoding"
)

type Server struct {
	host string
	port int
}

// Creates a new server
func NewServer(host string, port int) *Server {
	return &Server{host: host, port: port}
}

// Listens for incoming connections.
func (s *Server) Listen(callback func(net.Conn, encoding.Message)) error {
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

		go s.handle(conn, callback)
	}
}

// Handles incoming requests.
func (*Server) handle(conn net.Conn, callback func(net.Conn, encoding.Message)) error {
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	_, err = conn.Read(buf)
	if err != nil {
		return errors.New(fmt.Sprintf("Error reading: %s", err.Error()))
	}

	decoded, err := encoding.Unmarshal(string(buf))
	if err != nil {
		return err
	}

	go callback(conn, *decoded)

	return nil
}

// Sends an encoded message
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
