package main

import (
	"fmt"
	"net"
)

func main () {
	fmt.Println("HTTP SERVER RUNNING")
	ln, err := net.Listen("tcp", ":4221")
	if err != nil {
		fmt.Printf("Error binding to a port: %v\n", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error accepting TCP connection: %v\n", err)
		}
		
		fmt.Printf("New TCP connection accepted: %v\n", conn)

		n, err := conn.Read([]byte("GET / HTTP/1.1\r\n\r\n"))
		fmt.Printf("Reading from TCP connection: %v\n", n)

		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
		// go conn.Close()
	}
}
