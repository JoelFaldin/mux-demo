package models

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mux-demo/internal/protocol"
	"net"
	"sync"
)

type Stream struct {
	StreamID uint32
	Chan     chan []byte
	leftOver bytes.Buffer
	Conn     net.Conn
	once     sync.Once
}

// Used in server
// Offers an abstraction layer over Data,
// hiding the channels logic.
func (s *Stream) Read(p []byte) (n int, err error) {
	if s.leftOver.Len() != 0 {
		n, err = s.leftOver.Read(p)
		if err != nil {
			fmt.Println("[server] Error trying to read leftover", err.Error())
		}

		return n, nil
	}

	data, ok := <-s.Chan
	if !ok {
		return 0, io.EOF
	}

	_, err = s.leftOver.Write(data)
	if err != nil {
		fmt.Println("[server] Error writting data to leftover", err.Error())
		return
	}

	n, err = s.leftOver.Read(p)
	if err != nil && !errors.Is(err, io.EOF) {
		fmt.Println("[server] Error reading leftover buffer", err.Error())
		return
	}

	return n, nil
}

func (s *Stream) Write(b []byte) {
	protocol.WriteFrame(s.Conn, s.StreamID, protocol.FrameTypeNormal, b)
}

func (s *Stream) Close() {
	protocol.WriteFrame(s.Conn, s.StreamID, protocol.FrameTypeClose, []byte{})

	s.once.Do(func() {
		close(s.Chan)
	})
}
