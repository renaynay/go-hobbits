package tcp

import (
	"fmt"
	"net"
	"reflect"
	"strconv"
	"testing"

	"github.com/renaynay/go-hobbits/encoding"
)

func TestNewServer(t *testing.T) {
	var test = []struct {
		host   string
		port   int
		server *Server
	}{
		{host: "test", port: 3333, server: &Server{host: "test", port: 3333}},
		{host: "host", port: 4000, server: &Server{host: "host", port: 4000}},
	}

	for i, tt := range test {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			server := NewServer(tt.host, tt.port)
			if !reflect.DeepEqual(&server, &tt.server) {
				t.Errorf("return value of NewServer does not match expected value")
			}
		})
	}
}

func TestTCP(t *testing.T) {
	server := NewServer("", 1123)
	ch := make(chan encoding.Message)

	go server.Listen(func(_ net.Conn, message encoding.Message) {
		ch <- message
	})

	conn, err := net.Dial("tcp", ":1123")
	if err != nil {
		t.Error("could not connect to TCP server: ", err)
	}

	fmt.Println(conn)

	_, err = conn.Write([]byte("EWP 13.05 RPC blahblahblah json 16 14\nthis is a headerthis is a body"))
	if err != nil {
		t.Error("could not write to the TCP server: ", err)
	}
	read := <-ch

	expected := encoding.Message{
		Version:     "13.05",
		Protocol:    "RPC",
		Compression: "blahblahblah",
		Encoding:    "json",
		Headers:     []byte("this is a header"),
		Body:        []byte("this is a body"),
	}

	if !reflect.DeepEqual(expected, read) {
		t.Error("return value from TCP server does not match expected value")
	}
}
