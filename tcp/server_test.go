package tcp_test

import (
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/renaynay/go-hobbits/encoding"
	"github.com/renaynay/go-hobbits/tcp"
)

func TestTCP(t *testing.T) {
	server := tcp.NewServer("127.0.0.1", 0)
	ch := make(chan encoding.Message)

	go server.Listen(func(_ net.Conn, message encoding.Message) {
		ch <- message
	})

	for {
		if server.Addr() != nil {
			break
		}

		time.Sleep(1)
	}

	conn, err := net.Dial("tcp", server.Addr().String())
	if err != nil {
		t.Error("could not connect to TCP server: ", err)
	}

	_, err = conn.Write([]byte("EWP 13.05 RPC 16 14\nthis is a headerthis is a body"))
	if err != nil {
		t.Error("could not write to the TCP server: ", err)
	}
	read := <-ch

	expected := encoding.Message{
		Version:     "13.05",
		Protocol:    "RPC",
		Header:     []byte("this is a header"),
		Body:        []byte("this is a body"),
	}

	if !reflect.DeepEqual(expected, read) {
		t.Errorf("return value from TCP server does not match expected value. want=%v, got=%v", expected, read)
	}
}
