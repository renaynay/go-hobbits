// TODO: implement go fuzz
// TODO: implement mutation testing: https://github.com/zimmski/go-mutesting

package encoding

import (
	"encoding/binary"
	"errors"
)

// Unmarshal takes a wire protocol message and parses it
func Unmarshal(message []byte) (*Message, error) {
	var decoded Message

	if string(message[:3]) != "EWP" {
		return nil, errors.New("protocol unsupported, expecting 'EWP'")
	}

	decoded.Version = binary.BigEndian.Uint32(message[3:7])

	decoded.Protocol = Protocol(message[7])
	if decoded.Protocol > PING {
		return nil, errors.New("message protocol unsupported, expecting GOSSIP, RPC or PING")
	}

	headerLen := binary.BigEndian.Uint32(message[8:12])
	bodyLen := binary.BigEndian.Uint32(message[12:16])

	decoded.Header = message[16:16+headerLen]

	decoded.Body = message[16+headerLen:16+headerLen+bodyLen]


	return &decoded, nil
}
