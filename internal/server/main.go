package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mux-demo/internal/models"
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

		d := models.NewData()
		go func() {
			stream := d.GetFrame(1)
			buf := make([]byte, 512)
			for {
				n, err := stream.Read(buf)
				if err != nil {
					break
				}
				fmt.Println("Msg from channel:", string(buf[:n]))
			}
		}()
		for {
			h, res, err := utils.ReadFrame(conn)
			if err != nil && errors.Is(err, io.EOF) {
				break
			}

			stream := d.GetFrame(h.StreamID)

			stream.Chan <- res
		}
	}
}
