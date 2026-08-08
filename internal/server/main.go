package main

import (
	"fmt"
	"log"
	"mux-demo/internal/utils"
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

		h, res := utils.ReadFrame(conn)

		fmt.Println("Stream Id: ", h.StreamID)
		fmt.Println("Type: ", h.Type)
		fmt.Println("Data length: ", h.Length)
		fmt.Println("Res: ", res)
	}
}
