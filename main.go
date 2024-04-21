package main

import (
	"fmt"
	"net"
	"go-http-server/lib/utils/http"
)

func handleConnection(conn net.Conn) {
	readBuffer := make([]byte, 1024)
	n, err := conn.Read(readBuffer)

	if err != nil {
		fmt.Printf("Error reading from TCP connection: %v\n", err)
		return
	}

	connectionData := string(readBuffer[:n])
	fmt.Printf("\nReading from TCP connection:\n%v\n", connectionData)

	
	if http.GetRequestPath(connectionData) == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else {
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}

	defer conn.Close()
}

func main () {
	fmt.Println("HTTP SERVER RUNNING")

	ln, err := net.Listen("tcp", ":4221")
	if err != nil {
		fmt.Printf("Error binding to a port: %v\n", err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error accepting TCP connection: %v\n", err)
		}
		
		fmt.Printf("New TCP connection accepted")

		handleConnection(conn)
	}
}
