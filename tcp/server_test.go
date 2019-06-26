package tcp_test

import (
	"encoding/binary"
	"fmt"
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

	msg := getLength("EWP 13.05 RPC 16 14\nthis is a headerthis is a body")

	_, err = conn.Write(msg)
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

		ch <- string(read)
	}()

	msg := getLength("EWP 13.05 PING 4 14\npingthis is a body")

	_, err = conn.Write(msg)
	if err != nil {
		t.Error("could not write to the TCP server: ", err)
	}

	readFromCh := <-ch
	expected := "EWP 13.05 PING 4 14\npongthis is a body"

	if readFromCh != expected {
		t.Error("server does not send the correct default pong response to ping")
	}
}

func TestLongBody(t *testing.T) {
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

	longBody := "this is an extremely long body that is meant to be very very long asdsadofa asdpoifajpsdofijaspdof apofiajspdofajspdofjaspdofijaspdofjasdpofasjpdofiajspdfoiajspdofiajspdofiajspdofjaspdofjapsdoifjapsodfijpasodfijaspdofjaspdofapsdofiaspdofijapsdofpasdoifjapsodfjthis is an extremely long body that is meant to be very very long asdsadofa asdpoifajpsdofijaspdof apofiajspdofajspdofjaspdofijaspdofjasdpofasjpdofiajspdfoiajspdofiajspdofiajspdofjaspdofjapsdoifjapsodfijpasodfijaspdofjaspdofapsdofiaspdofijapsdofpasdoifjapsodfjthis is an extremely long body that is meant to be very very long asdsadofa asdpoifajpsdofijaspdof apofiajspdofajspdofjaspdofijaspdofjasdpofasjpdofiajspdfoiajspdofiajspdofiajspdofjaspdofjapsdoifjapsodfijpasodfijaspdofjaspdofapsdofiaspdofijapsdofpasdoifjapsodfjthis is an extremely long body that is meant to be very very long asdsadofa asdpoifajpsdofijaspdof apofiajspdofajspdofjaspdofijaspdofjasdpofasjpdofiajspdfoiajspdofiajspdofiajspdofjaspdofjapsdoifjapsodfijpasodfijaspdofjaspdofapsdofiaspdofijapsdofpasdoifjapsodfjthis is an extremely long body that is meant to be very very long asdsadofa asdpoifajpsdofijaspdof apofiajspdofajspdofjaspdofijaspdofjasdpofasjpdofiajspdfoiajspdofiajspdofiajspdofjaspdofjapsdoifjapsodfijpasodfijaspdofjaspdofapsdofiaspdofijapsdofpasdoifjapsodfjthis is an extremely long body that is meant to be very very long asdsadofa asdpoifajpsdofijaspdof apofiajspdofajspdofjaspdofijaspdofjasdpofasjpdofiajspdfoiajspdofiajspdofiajspdofjaspdofjapsdoifjapsodfijpasodfijaspdofjaspdofapsdofiaspdofijapsdofpasdoifjapsodfjthis is an extremely long body that is meant to be very very long asdsadofa asdpoifajpsdofijaspdof apofiajspdofajspdofjaspdofijaspdofjasdpofasjpdofiajspdfoiajspdofiajspdofiajspdofjaspdofjapsdoifjapsodfijpasodfijaspdofjaspdofapsdofiaspdofijapsdofpasdoifjapsodfjthis is an extremely long body that is meant to be very very long asdsadofa asdpoifajpsdofijaspdof apofiajspdofajspdofjaspdofijaspdofjasdpofasjpdofiajspdfoiajspdofiajspdofiajspdofjaspdofjapsdoifjapsodfijpasodfijaspdofjaspdofapsdofiaspdofijapsdofpasdoifjapsodfjthis is an extremely long body that is meant to be very very long asdsadofa asdpoifajpsdofijaspdof apofiajspdofajspdofjaspdofijaspdofjasdpofasjpdofiajspdfoiajspdofiajspdofiajspdofjaspdofjapsdoifjapsodfijpasodfijaspdofjaspdofapsdofiaspdofijapsdofpasdoifjapsodfjthis is an extremely long body that is meant to be very very long asdsadofa asdpoifajpsdofijaspdof apofiajspdofajspdofjaspdofijaspdofjasdpofasjpdofiajspdfoiajspdofiajspdofiajspdofjaspdofjapsdoifjapsodfijpasodfijaspdofjaspdofapsdofiaspdofijapsdofpasdoifjapsodfjthis is an extremely long body that is meant to be very very long asdsadofa asdpoifajpsdofijaspdof apofiajspdofajspdofjaspdofijaspdofjasdpofasjpdofijspdfoiajspdofiajspdofiajspdofjaspdofjapsdoifjapsodfijpasodfijaspdofjaspdofapsdofiaspdofijapsdofpasdoifjapsodfj"

	msg := getLength(fmt.Sprintf("EWP 13.05 RPC 16 2859\nthis is a header%s", longBody))

	_, err = conn.Write(msg)
	if err != nil {
		t.Error("could not write to the TCP server: ", err)
	}
	read := <-ch

	expected := encoding.Message{
		Version:  "13.05",
		Protocol: encoding.RPC,
		Header:   []byte("this is a header"),
		Body:     []byte(longBody),
	}

	if !reflect.DeepEqual(expected, read) {
		t.Errorf("return value from TCP server does not match expected value. want=%v, got=%v", expected, read)
	}
}

func getLength(message string) []byte {
	msg := []byte(message)

	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, uint32(len(msg)))

	return append(length, msg...)
}
