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

	utils.WriteFrame(conn, 1, 0, []byte("yo man"))
	utils.WriteFrame(conn, 2, 0, []byte("yo man 1"))
	utils.WriteFrame(conn, 1, 0, []byte("yo man 2"))

	conn.Close()
}
