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
			ch := d.GetFrame(1)
			for {
				msg := <-ch
				fmt.Println("Msg from channel:", string(msg))
			}
		}()
		for {
			h, res, err := utils.ReadFrame(conn)
			if err != nil && errors.Is(err, io.EOF) {
				break
			}

			// fmt.Println("Stream Id:", h.StreamID)
			// fmt.Println("Type:", h.Type)
			// fmt.Println("Data length:", h.Length)
			// fmt.Println("Res:", string(res))

			ch := d.GetFrame(h.StreamID)

			ch <- res
		}
	}
}
