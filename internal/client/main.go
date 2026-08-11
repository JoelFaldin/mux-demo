package main

import (
	"log"
	"mux-demo/internal/protocol"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal("[client] Couldnt connect to server", err.Error())
	}
	defer conn.Close()

	protocol.WriteFrame(conn, 1, 0, []byte("yo man"))
	protocol.WriteFrame(conn, 2, 0, []byte("yo man 1"))
	protocol.WriteFrame(conn, 1, 0, []byte("yo man 2"))
}
