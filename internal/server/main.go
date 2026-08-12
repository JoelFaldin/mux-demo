package main

import (
	"fmt"
	"log"
	"mux-demo/internal/models"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("[server] Error starting server", err.Error())
	}

	fmt.Println("[server] Server listening on port :8080!")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("[server] Error accepting request", err.Error())
		}

		fmt.Println("[server] Request accepted, processing...")

		s := models.NewSession(conn)

		go func() {
			stream := s.GetFrame(1, conn)
			buf := make([]byte, 512)
			for {
				n, err := stream.Read(buf)
				if err != nil {
					break
				}
				fmt.Println("Msg from channel:", string(buf[:n]))
			}
		}()
	}
}
