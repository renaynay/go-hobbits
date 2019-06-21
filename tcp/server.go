// Package tcp is a framework for a TCP server that will be able to handle TCP transport
// It will also decode inbound messages and encode outbound messages
package tcp

import (
	"fmt"
	"log"
	"net"

	"github.com/renaynay/go-hobbits/encoding"
)

// Callback is a function for message handling
type Callback func(net.Conn, encoding.Message)

type Server struct {
	host string
	port int

	addr net.Addr
}

// NewServer creates a new server
func NewServer(host string, port int) *Server {
	return &Server{host: host, port: port}
}

// Listen listens for incoming connections
func (s *Server) Listen(c Callback) error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		return fmt.Errorf("Error listening: %s", err.Error())
	}
	defer listen.Close()

	s.addr = listen.Addr()

	for {
		conn, err := listen.Accept()
		if err != nil {
			return fmt.Errorf(err.Error())
		}

		go func() {
			err := handle(conn, c)
			if err != nil {
				log.Print(err)
			}
		}()
	}
}

func (s Server) Addr() net.Addr {
	return s.addr
}

// handle handles incoming requests
func handle(conn net.Conn, c Callback) error {
	// TODO: find message size and store it in var. Wait til someone responds here: https://github.com/whiteblock/hobbits/issues/58#issuecomment-501883490
	// TODO: if no one responds, write a script to find the message size

	buf := make([]byte, 1024) // TODO: pass in message size here

	_, err := conn.Read(buf)
	if err != nil {
		return fmt.Errorf("Error reading: %s", err.Error())
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
