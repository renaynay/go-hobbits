# Hobbits

[![Build Status](https://travis-ci.com/renaynay/go-hobbits.svg?branch=master)](https://travis-ci.com/renaynay/go-hobbits) [![Go Report Card](https://goreportcard.com/badge/github.com/renaynay/go-hobbits)](https://goreportcard.com/report/github.com/renaynay/go-hobbits) [![API Reference](
                                                                                                                                                                                                                                                                             https://camo.githubusercontent.com/915b7be44ada53c290eb157634330494ebe3e30a/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f6c616e672f6764646f3f7374617475732e737667
                                                                                                                                                                                                                                                                             )](https://godoc.org/github.com/renaynay/go-hobbits)
                                                                                                                                                                                                                                                                             [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)                                                                                                                                                                                                                                                                             

Go implementation of [Hobbits](https://github.com/deltap2p/hobbits), a modular wire protocol.

## Installation

```
go get github.com/renaynay/go-hobbits 
```

## Usage

Encode a message: 

```go
var message encoding.Message{
	Version: "0.0",
	Protocol: "RPC",
	Header: []byte("test header"),
	Body: []byte("test body"),
}

encoded, err := encoding.Marshal(message)
if err != nil {
	log.Print(err)
}

fmt.Println(encoded)
```

Decode a message:

```go
decoded, err := encoding.Unmarshal("EWP 13.05 RPC 16 14\nthis is a headerthis is a body")
if err != nil {
	log.Print(err)
}

fmt.Println(decoded)
```

Here is a demo echo server: 

```go
server := tcp.NewServer("127.0.0.1", 1240)

err := server.Listen(func(conn net.Conn, message encoding.Message) {
    err := server.SendMessage(conn, message)
    if err != nil {
        log.Print(err)
    }
})

log.Print(err)
```