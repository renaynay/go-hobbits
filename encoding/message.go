// Package encoding implements message encoding and decoding for Hobbits, a Lightweight,
// Multiclient Wire Protocol For ETH2.0 Communications.
//
// By Rene Nayman
package encoding

type Protocol uint8

const (
	RPC    Protocol = iota
	GOSSIP
	PING
)

// Message represents a parsed Hobbits message.
// See examples of unparsed and parsed messages here: https://github.com/deltap2p/hobbits/blob/master/specs/protocol.md
type Message struct {
	Version  uint32
	Protocol Protocol
	Header   []byte
	Body     []byte
}
