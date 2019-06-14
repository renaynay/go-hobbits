package encoding

import (
	"errors"
	"fmt"
)

// Marshal takes a parsed message and encodes it to a wire protocol message
	func Marshal(message Message) (string, error) {
	if message.Version == "" {
		return "", errors.New("cannot marshal message, version not found")
	}

	if message.Protocol == "" {
		return "", errors.New("cannot marshal message, protocol not found")
	}

	return fmt.Sprintf(
		"EWP %s %s %d %d\n%s%s",
		message.Version,
		message.Protocol,
		len(string(message.Header)),
		len(string(message.Body)),
		string(message.Header),
		string(message.Body),
	), nil
}
