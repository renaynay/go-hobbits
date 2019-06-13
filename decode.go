package hobbits

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var AlphaNumRegex = regexp.MustCompile(`^[a-z0-9_]*$`)
var VersionNumRegex = regexp.MustCompile(`^(\d+\.)(\d+)*$`)

//TODO: check error messages

// Unmarshal takes a string and parses it to return a hobbit message
func Unmarshal(message string) (*Message, error) {
	var decoded Message

	lines := strings.Split(message, "\n")
	if len(lines) != 2 {
		return nil, errors.New("message request must contain 2 lines")
	}

	metadata := strings.Split(lines[0], " ")
	if len(metadata) != 7 {
		return nil, errors.New("not all metadata provided")
	}

	if !VersionNumRegex.MatchString(metadata[1]) {
		return nil, errors.New("EWP version cannot be parsed")
	}
	decoded.Version = metadata[1]

	if metadata[2] != "RPC" && metadata[2] != "GOSSIP" {
		return nil, errors.New("communication protocol unsupported")
	}
	decoded.Protocol = metadata[2]



	if !AlphaNumRegex.MatchString(metadata[3]) {
		return nil, errors.New("incorrect metadata format, cannot parse compression")
	}
	decoded.Compression = metadata[3]

	if !AlphaNumRegex.MatchString(metadata[4]) {
		return nil, errors.New("incorrect metadata format, cannot parse encoding")
	}
	decoded.Encoding = metadata[4]

	headLength, err := strconv.Atoi(metadata[5])
	if err != nil {
		return nil, errors.New("incorrect metadata format, cannot parse header-length")
	}
	decoded.Headers = []byte(lines[1][:headLength])

	bodyLength, err := strconv.Atoi(metadata[6])
	if err != nil {
		return nil, errors.New("incorrect metadata format, cannot parse body-length")
	}
	decoded.Body = []byte(lines[1][headLength:headLength+bodyLength])

	return &decoded, nil
}
