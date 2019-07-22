// Package tcp is a framework for a TCP server that will be able to handle TCP transport
// It will also decode inbound messages and encode outbound messages
package tcp

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/pkg/errors"
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
		return errors.Wrap(err, "error listening")
	}
	defer listen.Close()

	s.addr = listen.Addr()

	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}

		go func() {
			err := s.handle(conn, c)
			if err != nil {
				conn.Close()
				log.Print(err)
			}
		}()
	}
}

func (s Server) Addr() net.Addr {
	return s.addr
}

// handle handles incoming requests
func (s *Server) handle(conn net.Conn, c Callback) error {
	for {
		buf, err := Read(conn)
		if err != nil {
			rando := make([]byte, 1024)
			_, err := io.ReadFull(conn, rando)
			fmt.Println(rando) // TODO delete all this garbage

			return errors.Wrap(err, "error reading from conn")
		}

		decoded, err := encoding.Unmarshal(buf)
		if err != nil {
			return err
		}

		if decoded.Protocol == encoding.PING {
			decoded.Header = []byte("pong")

			err := s.SendMessage(conn, *decoded)
			if err != nil {
				return errors.Wrap(err, "PONG could not be sent")
			}

			continue
		}

		go c(conn, *decoded)
	}
}

// SendMessage sends an encoded message
func (*Server) SendMessage(conn net.Conn, message encoding.Message) error {
  // defer conn.Close()
  
	encoded := encoding.Marshal(message)

	_, err := conn.Write(encoded)
	if err != nil {
		return err
	}

	return nil
}

// Read reads a message from the connection
func Read(conn net.Conn) ([]byte, error) {
	metadata := make([]byte, 16)

	_, err := conn.Read(metadata)
	if err != nil {
		return nil, errors.Wrap(err, "error reading length")
	}

	headerLen := binary.BigEndian.Uint32(metadata[8:12])
	bodyLen := binary.BigEndian.Uint32(metadata[12:16])

	buf := make([]byte, (headerLen + bodyLen))

	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return nil, errors.Wrap(err, "error reading packet")
	}

	return append(metadata, buf...), nil
}
