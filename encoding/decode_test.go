package encoding_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/renaynay/go-hobbits/encoding"
)

func TestUnmarshal(t *testing.T) { // TODO: why the fuck isn't this working?
	var tests = []struct {
		message []byte
		decoded encoding.Message
	}{
		{
			message: []byte{69, 87, 80, 0, 0, 0, 3, 0, 0, 0, 0, 16, 0, 0, 0, 14, 116, 104, 105, 115, 32, 105, 115, 32, 97, 32, 104, 101, 97, 100, 101, 114, 116, 104, 105, 115, 32, 105, 115, 32, 97, 32, 98, 111, 100, 121},
			decoded: encoding.Message{
				Version: uint32(3),
				Protocol: encoding.RPC,
				Header: []byte("this is a header"),
				Body: []byte("this is a body"),
			},
		},
		{
			message: []byte{69, 87, 80, 0, 0, 0, 3, 1, 0, 0, 0, 16, 0, 0, 0, 14, 116, 104, 105, 115, 32, 105, 115, 32, 97, 32, 104, 101, 97, 100, 101, 114, 116, 104, 105, 115, 32, 105, 115, 32, 97, 32, 98, 111, 100, 121},
			decoded: encoding.Message{
				Version: uint32(3),
				Protocol: encoding.GOSSIP,
				Header: []byte("this is a header"),
				Body: []byte("this is a body"),
			},
		},
		{
			message: []byte{69, 87, 80, 0, 0, 0, 3, 2, 0, 0, 0, 4, 0, 0, 0, 0, 112, 105, 110, 103},
			decoded: encoding.Message{
				Version: uint32(3),
				Protocol: encoding.PING,
				Header: []byte("ping"),
				Body: []byte(""),
			},
		},
	}

	for i, tt := range tests{
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			unmarshaled, err := encoding.Unmarshal(tt.message)
			if err != nil {
				t.Error(err.Error())
			}

			if !reflect.DeepEqual(*unmarshaled, tt.decoded) {
				t.Error("return value of Unmarshal does not match expected value")
			}
		})
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		encoding.Unmarshal([]byte{69, 87, 80, 0, 0, 0, 3, 2, 0, 0, 0, 4, 0, 0, 0, 0, 112, 105, 110, 103})
	}
}
