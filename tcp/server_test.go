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

	msg := encoding.Marshal(encoding.Message{
		Version:  uint32(3),
		Protocol: encoding.RPC,
		Header:   []byte("this is a header"),
		Body:     []byte("this is a body"),
	})

	_, err = conn.Write(msg)
	if err != nil {
		t.Error("could not write to the TCP server: ", err)
	}
	read := <-ch

	expected := encoding.Message{
		Version:  uint32(3),
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
	ch := make(chan []byte)

	go server.Listen(func(_ net.Conn, message encoding.Message) {})

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
		read, err := tcp.Read(conn)
		if err != nil {
			t.Error(err)
		}

		ch <- read
	}()

	msg := encoding.Marshal(encoding.Message{
		Version:  uint32(3),
		Protocol: encoding.PING,
		Header:   []byte("ping"),
		Body:     []byte("body"),
	})

	_, err = conn.Write(msg)
	if err != nil {
		t.Error("could not write to the TCP server: ", err)
	}

	readFromCh := <-ch
	expected := encoding.Marshal(encoding.Message{
		Version:  uint32(3),
		Protocol: encoding.PING,
		Header:   []byte("pong"),
		Body:     []byte("body"),
	})

	if !reflect.DeepEqual(readFromCh, expected) {
		t.Error("server does not send the correct default pong response to ping")
	}
}
