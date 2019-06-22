package tcp_test

import (
	"io/ioutil"
	"log"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/renaynay/go-hobbits/encoding"
	"github.com/renaynay/go-hobbits/tcp"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	m.Run()
}

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
		Version:  "13.05",
		Protocol: encoding.RPC,
		Header:   []byte("this is a header"),
		Body:     []byte("this is a body"),
	}

	if !reflect.DeepEqual(expected, read) {
		t.Errorf("return value from TCP server does not match expected value. want=%v, got=%v", expected, read)
	}
}

func TestPING(t *testing.T) {
	server := tcp.NewServer("127.0.0.1", 0)
	ch := make(chan string)

	go server.Listen(func(_ net.Conn, message encoding.Message) {
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

	go func() {
		read, err := ioutil.ReadAll(conn)
		if err != nil {
			t.Error(err)
		}

		ch <- string(read)
	}()

	_, err = conn.Write([]byte("EWP 13.05 PING 4 14\npingthis is a body"))
	if err != nil {
		t.Error("could not write to the TCP server: ", err)
	}

	readFromCh := <-ch
	expected := "EWP 13.05 PING 4 14\npongthis is a body"

	if readFromCh != expected {
		t.Error("server does not send the correct default pong response to ping")
	}
}
