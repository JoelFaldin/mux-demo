package models

import (
	"bytes"
	"fmt"
	"io"
)

type Stream struct {
	StreamID uint32
	Chan     chan []byte
	LeftOver bytes.Buffer
}

// Used in server
// Offers an abstraction layer over Data,
// hiding the channels logic.
func (s *Stream) Read(p []byte) (n int, err error) {
	if s.LeftOver.Len() != 0 {
		n, err = s.LeftOver.Read(p)
		if err != nil {
			fmt.Println("[server] Error trying to read leftover", err.Error())
		}

		return n, nil
	}

	data, ok := <-s.Chan
	if !ok {
		return 0, io.EOF
	}

	_, err = s.LeftOver.Write(data)
	if err != nil {
		fmt.Println("[server] Error writting data to leftover", err.Error())
		return
	}

	n, err = s.LeftOver.Read(p)
	if err != nil {
		fmt.Println("[server] Error reading leftover buffer", err.Error())
		return
	}

	return n, nil
}
