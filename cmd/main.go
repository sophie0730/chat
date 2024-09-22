package main

import (
	"log"
	"net"

	"sophiehsu.com/chat/internal"
)

func main() {
	s := internal.NewServer()
	go s.Run()

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("start server on: 8000")

	//endless loop for accepting connection
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}

		go s.NewClient(conn)
	}
}
