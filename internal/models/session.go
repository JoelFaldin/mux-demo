package models

import (
	"errors"
	"io"
	"mux-demo/internal/protocol"
	"net"
	"sync"
)

type Session struct {
	Data        map[uint32]*Stream
	DataMutex   sync.Mutex
	AcceptChann chan *Stream
}

func NewSession(conn net.Conn) *Session {
	s := &Session{
		Data:        make(map[uint32]*Stream),
		AcceptChann: make(chan *Stream, 16),
	}

	go s.readLoop(conn)

	return s
}

// Used on NewSession function.
// Using a for loop, reads data from net.Conn.
// If the data's type is Close, closes the channel and notifies the other.
// If the data's type is of any other type, writes the data into the channel.
func (s *Session) readLoop(conn net.Conn) {
	for {
		h, res, err := protocol.ReadFrame(conn)
		if err != nil && errors.Is(err, io.EOF) {
			break
		}

		stream := s.GetFrame(h.StreamID, conn)
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
func (d *Session) GetFrame(streamId uint32, conn net.Conn) *Stream {
	d.DataMutex.Lock()
	defer d.DataMutex.Unlock()

	find, ok := d.Data[streamId]
	if !ok {
		s := &Stream{
			StreamID: streamId,
			Chan:     make(chan []byte, 10),
			Conn:     conn,
		}

		d.Data[streamId] = s
		d.AcceptChann <- s

		return s
	} else {
		return find
	}
}

func (s *Session) Accept() *Stream {
	stream := <-s.AcceptChann
	return stream
}
