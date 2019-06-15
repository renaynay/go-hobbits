package main

import (
	"log"
	"net"

	"github.com/renaynay/go-hobbits/encoding"
	"github.com/renaynay/go-hobbits/tcp"
)

func main() {
	server := tcp.NewServer("127.0.0.1", 1240)

	err := server.Listen(func(conn net.Conn, message encoding.Message) {
		err := server.SendMessage(conn, message)
		if err != nil {
			log.Print(err)
		}
	})

	log.Print(err)
}