package models

import (
	"errors"
	"io"
	"mux-demo/internal/protocol"
	"net"
	"sync"
)

type Session struct {
	Data      map[uint32]*Stream
	DataMutex sync.Mutex
}

func NewSession(conn net.Conn) *Session {
	s := &Session{
		Data: make(map[uint32]*Stream),
	}

	go s.ReadLoop(conn)

	return s
}

func (s *Session) ReadLoop(conn net.Conn) {
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
		return s
	} else {
		return find
	}
}
