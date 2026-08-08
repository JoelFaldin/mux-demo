package main

import (
	"log"
	"mux-demo/internal/utils"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal("[client] Couldnt connect to server", err.Error())
	}

	conn.Write([]byte("hi man\n"))

	utils.WriteFrame(conn, 1, 0, []byte("yo man"))
}
