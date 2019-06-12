package hobbits

import "testing"

func TestMarshal(t *testing.T) {
	var test = struct{
		message Message
	}{
		message: Message{
			Version: "13.05",
			Protocol: "RPC",
			Compression: "blahblahblah",
			Encoding: "JSON",
			Headers: []byte{123, 123, 123, 123},
			Body: []byte{456, 456, 789, 903, 65},
		},
	}
}
//TODO does encode need to find the header and body length too?