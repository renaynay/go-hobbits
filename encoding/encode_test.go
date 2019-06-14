package encoding

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func TestMarshal_Successful(t *testing.T) {
	var test = []struct {
		encoded Message
		message string
	}{
		{
			encoded: Message{
				Version:     "13.05",
				Protocol:    "RPC",
				Header:     []byte("this is a header"),
				Body:        []byte("this is a body"),
			},
			message: "EWP 13.05 RPC 16 14\nthis is a headerthis is a body",
		},
		{
			encoded: Message{
				Version:     "13.05",
				Protocol:    "GOSSIP",
				Header:     []byte("testing"),
				Body:        []byte("testing body"),
			},
			message: "EWP 13.05 GOSSIP 7 12\ntestingtesting body",
		},
		{
			encoded: Message{
				Version:     "1230329483.05392489",
				Protocol:    "RPC",
				Header:     []byte("test"),
				Body:        []byte("test"),
			},
			message: "EWP 1230329483.05392489 RPC 4 4\ntesttest",
		},
	}

	for i, tt := range test {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			string, _ := Marshal(tt.encoded)
			if !reflect.DeepEqual(string, tt.message) {
				t.Errorf("return value of Marshal did not match expected value. wanted: %v, got: %v", tt.message, string)
			}
		})
	}
}

func TestMarshal_Unsuccessful(t *testing.T) {
	var test = []struct {
		encoded Message
		err     error
	}{
		{
			encoded: Message{
				Version:     "",
				Protocol:    "RPC",
				Header:     []byte("this is a header"),
				Body:        []byte("this is a body"),
			},
			err: errors.New("cannot marshal message, version not found"),
		},
		{
			encoded: Message{
				Version:     "1230329483.05392489",
				Protocol:    "",
				Header:     []byte("test"),
				Body:        []byte("test"),
			},
			err: errors.New("cannot marshal message, protocol not found"),
		},
	}

	for i, tt := range test {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			_, err := Marshal(tt.encoded)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("return value of Marshal did not match expected value")
			}
		})
	}
}
