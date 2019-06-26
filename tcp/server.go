// Package tcp is a framework for a TCP server that will be able to handle TCP transport
// It will also decode inbound messages and encode outbound messages
package tcp

import (
	"fmt"
	"io"
	"log"
	"net"
	"encoding/binary"

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
	pktLen := make([]byte, 4)

	_, err := conn.Read(pktLen)
	if err != nil {
		return errors.Wrap(err, "error reading length")
	}

	packetLength := binary.BigEndian.Uint32(pktLen)

	buf := make([]byte, packetLength)

	_, err = io.ReadFull(conn, buf)
	if err != nil {
		return errors.Wrap(err, "error reading packet")
	}

	decoded, err := encoding.Unmarshal(string(buf))
	if err != nil {
		return err
	}

	if decoded.Protocol == "PING" {
		decoded.Header = []byte("pong")

		err := s.SendMessage(conn, *decoded)
		if err != nil {
			return errors.Wrap(err, "PONG could not be sent")
		}

		return nil
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

	packet := []byte(encoded)
	packetLength := len(encoded)
	pktLen := make([]byte, 4)

	binary.BigEndian.PutUint32(pktLen[0:], uint32(packetLength))

	pktLen = append(pktLen, packet...)

	_, err = conn.Write([]byte(pktLen))
	if err != nil {
		return err
	}

	return nil
}
