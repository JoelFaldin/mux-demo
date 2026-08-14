# mux-demo - String multiplexing over a TCP connection

Bridge project built using [Go](https://go.dev/) to understand how tools like ngrok handle different logical connections (_streams_) over a unique TCP connection. It implements a (very simple) custom binary protocol with framing based on headers of fixed length, plus a concurrent layer with goroutines and channels to _demultiplex_ incoming traffic.

## Project Structure

```
tcp-mux/
├── protocol/
│   ├── header.go         # struct Header (StreamID, Type, Length) + frameType (iota)
│   ├── frames.go         # Functions that prepare to talk with server and client
│
├── models/
│   ├── stream.go         # struct Stream: Read/Write/Close,
│   └── session.go        # struct Session: stream map + mutex, ReadLoop, Accept, OpenStream
│
├── cmd/
│   ├── server/
│   │   └── main.go       # Listens, accept connections, create Session, loops over Accept()
│   └── client/
│       └── main.go       # Connects, create Session, open streams with OpenStream()
│
└── go.mod
```

## The Protocol

Each message that goes through the connection carries a _header_ of 9 bytes before payload:

```
[ StreamID (4 bytes) ][ Type (1 byte) ][ Length (4 bytes) ][ ...payload... ]
```

* *StreamID*: Which conversation does the frame belong to
* *Type*: What kind of connection is being used (`FrameTypeNormal`, `FrameTypeClose`)
* *Length*: The amount of payload bytes are there, to know exactly where does the frame end

Everyting is encoded in bigendian - the standard convention in network protocols (network byte order).

## Prerequisites

* Golang 1.21 or higher.
* (Optional to inspect network traffic) Wireshark.

## Instalation

1. Clone the repo.
2. `go mod tidy`

## Usage

1. Start the server:

```
go run internal/cmd/server/main.go
```

2. In a different terminal, run the client:

```
go run internal/cmd/client/main.go
```

## Thanks for visiting!
