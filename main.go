package main

import (
	"fmt"
	"net"
)

func main () {
	fmt.Println("HTTP SERVER RUNNING")
	ln, err := net.Listen("tcp", ":4221")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(conn)
	}
}
