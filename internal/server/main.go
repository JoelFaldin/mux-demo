package main

import (
	"bufio"
	"fmt"
	"log"
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

		reader := bufio.NewReader(conn)
		res, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("[server] Couldnt read client data", err.Error())
		}

		fmt.Println(res)
	}

}
