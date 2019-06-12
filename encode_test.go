package hobbits

import (
	"testing"
	"strconv"
	"reflect"
)

func TestMarshal(t *testing.T) {
	var test = []struct {
		encoded Message
		message string
	}{
		{
			encoded: Message{
				Version: "13.05",
				Protocol: "RPC",
				Compression: "blahblahblah",
				Encoding: "JSON",
				Headers: []byte("this is a header"),
				Body: []byte("this is a body"),
			},
			message: "EWP 13.05 RPC blahblahblah JSON 16 14\nthis is a headerthis is a body",
		},
		{
			encoded: Message{
				Version: "13.05",
				Protocol: "GOSSIP",
				Compression: "blahblahb123_f",
				Encoding: "BSON",
				Headers: []byte("testing"),
				Body: []byte("testing body"),
			},
			message: "EWP 13.05 GOSSIP blahblahb123_f BSON 7 12\ntestingtesting body",
		},
		{
			encoded: Message{
				Version: "1230329483.05392489",
				Protocol: "RPC",
				Compression: "blahblahblah",
				Encoding: "JSON",
				Headers: []byte("test"),
				Body: []byte("test"),
			},
			message: "EWP 1230329483.05392489 RPC blahblahblah JSON 4 4\ntesttest",
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
