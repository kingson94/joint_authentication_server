package main

import (
	"log"
	"net"
)

func main() {
	log.Println("First program")
	server, error := net.Listen("tcp4", ":8088")
	if error != nil {
		log.Println(error.Error())
	}
	for {
		server.Accept()
	}
	server.Close()
}
