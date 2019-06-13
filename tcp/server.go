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
func Listen(connType string, connHost string, connPort string) error {
	listen, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		return errors.New(fmt.Sprintf("Error listening: %s", (err.Error())))
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}

		go handleRequest(conn)
	}

	return nil
}

// Handles incoming requests.
func handleRequest(conn net.Conn) error {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		return errors.New(fmt.Sprintf("Error reading: %s", err.Error()))
	}

	// Decodes data
	_, err = encoding.Unmarshal(string(buf)) // TODO: what do I do with the unmarshaled message? should it be left blank bc this is a framework ?
	if err != nil {
		return err
	}

	return nil
}

// Sends an encoded message
func SendMessage(conn net.Conn, message encoding.Message) error { // TODO: where does this get called?
	// Encodes the message
	encoded, err := encoding.Marshal(message)
	if err != nil {
		return err
	}

	// Send a response back to person contacting us.
	conn.Write([]byte(encoded))
	// Close the connection when you're done with it.
	conn.Close()

	return nil
}
