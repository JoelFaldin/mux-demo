package utils

import (
	"encoding/binary"
	"fmt"
	"net"
)

func WriteFrame(conn net.Conn, streamId uint32, t_type uint32, payload []byte) {
	finalSlice := make([]byte, 9+len(payload))

	binary.BigEndian.PutUint32(finalSlice, streamId)
	finalSlice[4] = byte(t_type)
	binary.BigEndian.PutUint32(finalSlice[5:9], uint32(len(payload)))

	copy(finalSlice[9:], payload)

	fmt.Printf("%v\n", finalSlice)
	fmt.Printf("%x\n", finalSlice)
}
