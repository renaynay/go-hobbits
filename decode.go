package hobbits

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

//TODO: unmarshal takes a string and parses it to return a hobbit message
//TODO: check error messages

//TODO: DOCUMENT THIS SHIT
func Unmarshal(req string) (Message, error) {
	var decoded Message

	lines := strings.Split(req, "\n")
	if len(lines) != 2 {
		return Message{}, errors.New("message request must contain 2 lines")
	}

	metadata := strings.Split(lines[0], " ")
	if len(metadata) != 7 {
		return Message{}, errors.New("not enough metadata for parsing")
	}

	if !regexp.MustCompile(`^(\d+\.)(\d+)*$`).MatchString(metadata[1]) {
		return Message{}, errors.New("EWP version cannot be parsed")
	}
	decoded.Version = metadata[1]

	if metadata[2] != "RPC" && metadata[2] != "GOSSIP" {
		return Message{}, errors.New("communication protocol unsupported")
	}
	decoded.Protocol = metadata[2]

	if !regexp.MustCompile(`^(\d+\.)(\d+)*$`).MatchString(metadata[3]) {
		return Message{}, errors.New("incorrect metadata format, cannot parse compression")
	}
	decoded.Compression = metadata[3]

	if !regexp.MustCompile(`^(\d+\.)(\d+)*$`).MatchString(metadata[4]) {
		return Message{}, errors.New("incorrect metadata format, cannot parse encoding")
	}
	decoded.Encoding = metadata[4]

	headLength, err := strconv.Atoi(metadata[4])
	if err != nil {
		// throw an error?
	}
	decoded.Headers = []byte(metadata[6][:headLength]) // does this work how i want it to?

	bodyLength, err := strconv.Atoi(metadata[5]) // do we even need this?
	if err != nil {
		// throw an error
	}
	if bodyLength > 0 {
		decoded.Body = []byte(metadata[6][headLength:])
	}

	return decoded, nil
}
