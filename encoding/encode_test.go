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
				Compression: "blahblahblah",
				Encoding:    "json",
				Headers:     []byte("this is a header"),
				Body:        []byte("this is a body"),
			},
			message: "EWP 13.05 RPC blahblahblah json 16 14\nthis is a headerthis is a body",
		},
		{
			encoded: Message{
				Version:     "13.05",
				Protocol:    "GOSSIP",
				Compression: "blahblahb123_f",
				Encoding:    "bson",
				Headers:     []byte("testing"),
				Body:        []byte("testing body"),
			},
			message: "EWP 13.05 GOSSIP blahblahb123_f bson 7 12\ntestingtesting body",
		},
		{
			encoded: Message{
				Version:     "1230329483.05392489",
				Protocol:    "RPC",
				Compression: "blahblahblah",
				Encoding:    "json",
				Headers:     []byte("test"),
				Body:        []byte("test"),
			},
			message: "EWP 1230329483.05392489 RPC blahblahblah json 4 4\ntesttest",
		},
	}

	for i, tt := range test {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			string, _ := Marshal(tt.encoded)
			if !reflect.DeepEqual(string, tt.message) {
				t.Errorf("return value of Marshal did not match expected value")
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
				Compression: "blahblahblah",
				Encoding:    "json",
				Headers:     []byte("this is a header"),
				Body:        []byte("this is a body"),
			},
			err: errors.New("cannot marshal message, version not found"),
		},
		{
			encoded: Message{
				Version:     "13.05",
				Protocol:    "GOSSIP",
				Compression: "",
				Encoding:    "bson",
				Headers:     []byte("testing"),
				Body:        []byte("testing body"),
			},
			err: errors.New("cannot marshal message, compression not found"),
		},
		{
			encoded: Message{
				Version:     "1230329483.05392489",
				Protocol:    "RPC",
				Compression: "blahblahblah",
				Encoding:    "",
				Headers:     []byte("test"),
				Body:        []byte("test"),
			},
			err: errors.New("cannot marshal message, encoding not found"),
		},
		{
			encoded: Message{
				Version:     "1230329483.05392489",
				Protocol:    "",
				Compression: "blahblahblah",
				Encoding:    "json",
				Headers:     []byte("test"),
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
