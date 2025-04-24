package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting ISO-8583 Authorizer...")
	// Listen on TCP port 9000
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	defer ln.Close()
	fmt.Println("Listening on port 9000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	// Simple response for testing
	fmt.Fprintln(conn, "ISO-8583 Authorizer POC")
}
