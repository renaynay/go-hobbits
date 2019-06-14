
// TODO: implement go fuzz
// TODO: implement mutation testing: https://github.com/zimmski/go-mutesting

package encoding

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var alphaNumRegex = regexp.MustCompile(`^[a-z0-9_]*$`)
var versionNumRegex = regexp.MustCompile(`^(\d+\.)(\d+)*$`)

// Unmarshal takes a wire protocol message and parses it
func Unmarshal(message string) (*Message, error) {
	var decoded Message

	lines := strings.Split(message, "\n")
	if len(lines) != 2 {
		return nil, errors.New("message request must contain 2 lines")
	}

	metadata := strings.Split(lines[0], " ")
	if len(metadata) != 5 {
		return nil, errors.New("not all metadata provided")
	}

	if !versionNumRegex.MatchString(metadata[1]) {
		return nil, errors.New("EWP version cannot be parsed")
	}
	decoded.Version = metadata[1]

	if metadata[2] != "RPC" && metadata[2] != "GOSSIP" {
		return nil, errors.New("communication protocol unsupported")
	}
	decoded.Protocol = metadata[2]

	headLength, err := strconv.Atoi(metadata[5])
	if err != nil {
		return nil, errors.New("incorrect metadata format, cannot parse header-length")
	}
	decoded.Header = []byte(lines[1][:headLength])

	bodyLength, err := strconv.Atoi(metadata[6])
	if err != nil {
		return nil, errors.New("incorrect metadata format, cannot parse body-length")
	}
	decoded.Body = []byte(lines[1][headLength:headLength+bodyLength])

	return &decoded, nil
}
