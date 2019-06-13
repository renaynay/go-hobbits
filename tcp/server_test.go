package tcp

import (
	"reflect"
	"strconv"
	"testing"
)

func TestNewServer(t *testing.T) {
	var test = []struct {
		host string
		port int
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

// TODO: this isn't a unit-test type test, it's more of an integration test
// TODO:
//  1) pipe a channel through handle (do this by creating a chan in callback?)
//  2) ignore conn



//func TestServer_Listen(t *testing.T) {
//	//TODO: complete test
//}
//
//func Test_handle(t *testing.T) {
//	var test = []struct {
//		conn net.Conn
//		c Callback
//	}{
//		{}
//	}
//	//TODO: complete test
//}
//
//func TestServer_SendMessage(t *testing.T) {
//	//TODO: complete test
//}
