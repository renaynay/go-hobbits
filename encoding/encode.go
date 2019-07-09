package encoding

import (
	"encoding/binary"
)

// Marshal takes a parsed message and encodes it to a wire protocol message
func Marshal(message Message) []byte {
	marshaled := []byte{}

	marshaled = append(marshaled, []byte("EWP")...)

	version := make([]byte, 4)
	binary.BigEndian.PutUint32(version, message.Version)

	marshaled = append(marshaled, version...)
	marshaled = append(marshaled, byte(message.Protocol))

	head := uint32(len(message.Header))
	headerLen := make([]byte, 4)
	binary.BigEndian.PutUint32(headerLen, head)

	marshaled = append(marshaled, headerLen...)

	body := uint32(len(message.Body))
	bodyLen := make([]byte, 4)
	binary.BigEndian.PutUint32(bodyLen, body)

	marshaled = append(marshaled, bodyLen...)

	marshaled = append(marshaled, message.Header...)
	marshaled = append(marshaled, message.Body...)

	return marshaled
}
