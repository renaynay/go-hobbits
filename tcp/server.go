//TODO: godocs
package tcp

import (
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/renaynay/go-hobbits/encoding"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

//TODO: shouldn't be in a main func

// Listens for incoming connections.
func listen() error {
	listen, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		return errors.New(fmt.Sprintf("Error listening: %s", (err.Error())))
		os.Exit(1)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
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

	_, err = encoding.Unmarshal(string(buf)) // TODO: what do I do with the unmarshaled message? should it be left blank bc this is a framework ?
	if err != nil {
		fmt.Println(err) // TODO: do i need to print?
		return err
	}

	return nil
}

func sendMessage(conn net.Conn, message encoding.Message) error { // TODO: where does this get called?
	encoded, err := encoding.Marshal(message)
	if err != nil {
		fmt.Println(err) // TODO do i need to print?
		return err
	}

	// Send a response back to person contacting us.
	conn.Write([]byte(encoded))
	// Close the connection when you're done with it.
	conn.Close()

	return nil
}