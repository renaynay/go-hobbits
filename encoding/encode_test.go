package encoding_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/renaynay/go-hobbits/encoding"
)

func TestMarshal_Successful(t *testing.T) {
	var tests = []struct {
		encoded encoding.Message
		message []byte
	}{
		{
			encoded: encoding.Message{
				Version: uint32(3),
				Protocol: encoding.RPC,
				Header: []byte("this is a header"),
				Body: []byte("this is a body"),
			},
			message: []byte{69, 87, 80, 0, 0, 0, 0, 16, 0, 0, 0, 14, 116, 104, 105, 115, 32, 105, 115, 32, 97, 32, 104, 101, 97, 100, 101, 114, 116, 104, 105, 115, 32, 105, 115, 32, 97, 32, 98, 111, 100, 121},
		},
		{
			encoded: encoding.Message{
				Version: uint32(3),
				Protocol: encoding.GOSSIP,
				Header: []byte("this is a header"),
				Body: []byte("this is a body"),
			},
			message: []byte{69, 87, 80, 1, 0, 0, 0, 16, 0, 0, 0, 14, 116, 104, 105, 115, 32, 105, 115, 32, 97, 32, 104, 101, 97, 100, 101, 114, 116, 104, 105, 115, 32, 105, 115, 32, 97, 32, 98, 111, 100, 121},
		},
		{
			encoded: encoding.Message{
				Version: uint32(3),
				Protocol: encoding.PING,
				Header: []byte("ping"),
				Body: []byte(""),
			},
			message: []byte{69, 87, 80, 2, 0, 0, 0, 4, 0, 0, 0, 0, 112, 105, 110, 103},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if !reflect.DeepEqual(encoding.Marshal(tt.encoded), tt.message) {
				t.Error("return value of Marshal does not match expected value")
			}
		})
	}
}

func BenchmarkMarshal(b *testing.B) {

}