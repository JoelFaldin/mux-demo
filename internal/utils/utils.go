package utils

import (
	"encoding/binary"
	"io"
	"log"
	"mux-demo/internal/headers"
	"net"
)

// Used in the client
// Given streamId, type and payload, WriteFrame builds a complete slice
// with data to send to backend.
func WriteFrame(conn net.Conn, streamId uint32, t_type uint32, payload []byte) {
	finalSlice := make([]byte, 9+len(payload))

	binary.BigEndian.PutUint32(finalSlice, streamId)
	finalSlice[4] = byte(t_type)
	binary.BigEndian.PutUint32(finalSlice[5:9], uint32(len(payload)))

	copy(finalSlice[9:], payload)

	conn.Write(finalSlice)
}

// Used in the server
// Prepares slices to read from net.Conn, and parses header
// to extract data
func ReadFrame(conn net.Conn) (headers.Header, string) {
	header := make([]byte, 9)

	_, err := io.ReadFull(conn, header)
	if err != nil {
		log.Println("[server] Error reading client data", err.Error())
	}

	streamId := binary.BigEndian.Uint32(header[0:4])
	t_type := header[4]

	data_length := binary.BigEndian.Uint32(header[5:])

	buf := make([]byte, data_length)
	r, err := io.ReadFull(conn, buf)
	if err != nil {
		log.Println("[server] Couldnt read client data", err.Error())
		return headers.Header{}, ""
	}

	receivedData := buf[:r]

	data_slice := make([]byte, data_length)
	copy(data_slice, receivedData)

	header_data := headers.Header{
		StreamID: streamId,
		Type:     t_type,
		Length:   data_length,
	}

	res := string(data_slice)

	return header_data, res
}
