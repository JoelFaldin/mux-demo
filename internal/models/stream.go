package models

import "io"

type Stream struct {
	StreamID uint32
	Chan     chan []byte
}

// Used in server
// Offers an abstraction layer over Data,
// hiding the channels logic.
func (s *Stream) Read(p []byte) (n int, err error) {
	data, ok := <-s.Chan
	if !ok {
		return 0, io.EOF
	}

	n = copy(p, data)
	return n, nil
}
