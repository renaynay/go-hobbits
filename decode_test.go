package hobbits

import (
	"reflect"
	"strconv"
	"testing"
)

func TestUnmarshal_Successful(t *testing.T) {
	var test = []struct {
		message string
		output Message
	}{
		{
			message: "EWP 13.05 RPC blahblahblah json 16 14\nthis is a headerthis is a body",
			output: Message{
				Version: "13.05",
				Protocol: "RPC",
				Compression: "blahblahblah",
				Encoding: "json",
				Headers: []byte("this is a header"),
				Body: []byte("this is a body"),
			},
		},
		{
			message: "EWP 13.05 GOSSIP blahblahb123_f bson 7 12\ntestingtesting body",
			output: Message{
				Version: "13.05",
				Protocol: "GOSSIP",
				Compression: "blahblahb123_f",
				Encoding: "bson",
				Headers: []byte("testing"),
				Body: []byte("testing body"),
			},
		},
		{
			message: "EWP 1230329483.05392489 RPC blahblahblah json 4 4\ntesttest",
			output: Message{
				Version: "1230329483.05392489",
				Protocol: "RPC",
				Compression: "blahblahblah",
				Encoding: "json",
				Headers: []byte("test"),
				Body: []byte("test"),
			},
		},
	}

	for i, tt := range test {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			output, _ := Unmarshal(tt.message)
			if !reflect.DeepEqual(output, tt.output) {
				t.Errorf("return value of Unmarshal does not match expected value")
			}
		})
	}
}
