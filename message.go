package hobbits

type Message struct {
	Version string
	Protocol string
	Compression string
	Encoding string
	Headers []byte
	Body []byte
}
