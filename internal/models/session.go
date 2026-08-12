package models

import (
	"errors"
	"io"
	"mux-demo/internal/protocol"
	"net"
	"sync"
	"sync/atomic"
)

type Session struct {
	Data        map[uint32]*Stream
	DataMutex   sync.Mutex
	AcceptChann chan *Stream
	nextId      atomic.Uint32
	Conn        net.Conn
}

func NewSession(conn net.Conn) *Session {
	s := &Session{
		Data:        make(map[uint32]*Stream),
		AcceptChann: make(chan *Stream, 16),
		Conn:        conn,
	}

	go s.readLoop()

	return s
}

// Used on NewSession function.
// Using a for loop, reads data from net.Conn.
// If the data's type is Close, closes the channel and notifies the other.
// If the data's type is of any other type, writes the data into the channel.
func (s *Session) readLoop() {
	for {
		h, res, err := protocol.ReadFrame(s.Conn)
		if err != nil && errors.Is(err, io.EOF) {
			break
		}

		stream := s.GetFrame(h.StreamID)
		if h.Type == byte(protocol.FrameTypeClose) {
			stream.Close()
			delete(s.Data, h.StreamID)
		} else {
			stream.Chan <- res
		}
	}
}

// Used in server
// Given streamId, searchs in models.Data map
// if theres any existent frame with that id.
func (d *Session) GetFrame(streamId uint32) *Stream {
	d.DataMutex.Lock()
	defer d.DataMutex.Unlock()

	find, ok := d.Data[streamId]
	if !ok {
		s := d.newAndRegisterStream(streamId)

		return s
	} else {
		return find
	}
}

// Used in server, for now in a test goroutine.
// Waits for any incoming connection. Doesn't need to know StreamId beforehand.
// Returns *Stream.
func (s *Session) Accept() *Stream {
	stream := <-s.AcceptChann
	return stream
}

func (s *Session) OpenStream() *Stream {
	newStreamId := s.nextId.Add(1)

	d := s.newAndRegisterStream(uint32(newStreamId))

	return d
}

func (s *Session) newAndRegisterStream(strId uint32) *Stream {
	ses := &Stream{
		StreamID: strId,
		Chan:     make(chan []byte, 10),
		Conn:     s.Conn,
	}

	s.Data[strId] = ses
	s.AcceptChann <- ses

	return ses
}
