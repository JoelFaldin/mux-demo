package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal("[client] Couldnt connect to server", err.Error())
	}

	conn.Write([]byte("hi man\n"))

	time.Sleep(500 * time.Millisecond)
}
