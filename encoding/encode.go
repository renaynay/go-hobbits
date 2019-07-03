package encoding

import (
	"encoding/binary"
)

// Marshal takes a parsed message and encodes it to a wire protocol message
func Marshal(message Message) []byte {
	marshaled := []byte{}

	marshaled = append(marshaled, []byte("EWP")...)
	marshaled = append(marshaled, byte(message.Protocol))

	head := uint32(len(message.Header))
	headerLen := []byte{}

	binary.BigEndian.PutUint32(headerLen, head)

	body := uint32(len(message.Body))
	bodyLen := []byte{}

	binary.BigEndian.PutUint32(bodyLen, body)

	marshaled = append(marshaled, headerLen...)
	marshaled = append(marshaled, bodyLen...)

	marshaled = append(marshaled, message.Header...)
	marshaled = append(marshaled, message.Body...)

	return marshaled
}
