package hobbits

import (
	"errors"
	"fmt"
)

// Marshal takes a hobbit message and encodes it to a string
func Marshal(message Message) (string, error) {
	if message.Version == "" {
		return "", errors.New("cannot marshal message, version not found")
	}

	if message.Protocol == "" {
		return "", errors.New("cannot marshal message, protocol not found")
	}

	if message.Compression == "" {
		return "", errors.New("cannot marshal message, compression not found")
	}

	if message.Encoding == "" {
		return "", errors.New("cannot marshal message, encoding not found")
	}

	return fmt.Sprintf(
		"EWP %s %s %s %s %d %d\n%s%s",
		message.Version,
		message.Protocol,
		message.Compression,
		message.Encoding,
		len(string(message.Headers)),
		len(string(message.Body)),
		string(message.Headers),
		string(message.Body),
	), nil
}
