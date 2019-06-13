// Package tcp is a framework for a TCP server that will be able to handle TCP transport
// It will also decode inbound messages and encode outbound messages
package tcp

import (
	"errors"
	"fmt"
	"net"

	"github.com/renaynay/go-hobbits/encoding"
)

// Listens for incoming connections.
func Listen(host string, port int, callback func(net.Conn, encoding.Message)) error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return errors.New(fmt.Sprintf("Error listening: %s", (err.Error())))
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}

		go handleRequest(conn, callback)
	}

	return nil
}

// Handles incoming requests.
func handleRequest(conn net.Conn, callback func(net.Conn, encoding.Message)) error {
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		return errors.New(fmt.Sprintf("Error reading: %s", err.Error()))
	}

	decoded, err := encoding.Unmarshal(string(buf)) // TODO: what do I do with the unmarshaled message? should it be left blank bc this is a framework ?
	if err != nil {
		return err
	}

	go callback(conn, *decoded)

	return nil
}

// Sends an encoded message
func SendMessage(conn net.Conn, message encoding.Message) error { // TODO: where does this get called?
	encoded, err := encoding.Marshal(message)
	if err != nil {
		return err
	}

	conn.Write([]byte(encoded))
	conn.Close()

	return nil
}
