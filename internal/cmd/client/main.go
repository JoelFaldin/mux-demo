package main

import (
	"log"
	"mux-demo/internal/models"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal("[client] Couldnt connect to server", err.Error())
	}
	defer conn.Close()

	session := models.NewSession(conn)
	stream := session.OpenStream()

	stream.Write([]byte("algo 1"))
}
