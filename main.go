package main

import (
	"fmt"
	"net"
	"go-http-server/lib/utils/http"
	"regexp"
	"strings"
	"flag"
	"os"
	"path/filepath"
)

func handleConnection(conn net.Conn, directory string) {
	readBuffer := make([]byte, 1024)
	n, err := conn.Read(readBuffer)

	if err != nil {
		fmt.Printf("Error reading from TCP connection: %v\n", err)
		return
	}

	connectionData := string(readBuffer[:n])
	fmt.Printf("\nReading from TCP connection:\n%v\n", connectionData)

	pattern := "/echo/.+"

	echoRegExp, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex pattern:", err)
		return
	}

	requestpath := http.GetRequestPath(connectionData)

	if requestpath == "/" {
		body := "server responded to /"
		response := http.BuildResponse(
			200, 
			"OK", 
			body, 
			"Content-Type: text/plain", 
			fmt.Sprintf("Content-Length: %v", len(body)),
		)

		conn.Write(response)
	} else if strings.Contains(connectionData, "POST /files") {
		fileName :=  strings.Split(requestpath, "/")[2]

		fullPath := filepath.Join(directory, fileName)

		reqBody := strings.Split(connectionData, "\r\n\r\n")[1]

		file, err := os.Create(fullPath)  //create a new file
		if err != nil {
				fmt.Println(err)
				return
		}
    
		defer file.Close()

		writeErr := os.WriteFile(fullPath, []byte(reqBody), 0666)
		if writeErr != nil {
			fmt.Println(writeErr)
		}

		body := reqBody 
		response := http.BuildResponse(
			201,
			"OK", 
			body, 
			"Content-Type: application/octet-stream", 
			fmt.Sprintf("Content-Length: %v", len(body)),
		)

		conn.Write(response)
	} else if strings.Contains(requestpath, "/files") {
		fileName :=  strings.Split(requestpath, "/")[2]

		fullPath := filepath.Join(directory, fileName)

		_, err := os.Stat(fullPath)
		if err != nil {
			fmt.Println("Specified file does not exist")
			body := "Specified file does not exist" 
			response := http.BuildResponse(
				404,
				"Not Found", 
				body, 
				"Content-Type: text/plain", 
				fmt.Sprintf("Content-Length: %v", len(body)),
			)
		conn.Write(response)
		return
		}

		data, err := os.ReadFile(fullPath)
		fileData := string(data)
		fmt.Printf("file data: %v\n", fileData)

		body := fileData 
		response := http.BuildResponse(
			200,
			"OK", 
			body, 
			"Content-Type: application/octet-stream", 
			fmt.Sprintf("Content-Length: %v", len(body)),
		)

		conn.Write(response)
	} else if requestpath == "/user-agent" {
		var body string
		for _, str := range strings.Split(connectionData, "\r\n") {
			if strings.Contains(str, "User-Agent:") {
				body = strings.Split(str, " ")[1]
			}
		}
		response := http.BuildResponse(
			200,
			"OK", 
			body, 
			"Content-Type: text/plain", 
			fmt.Sprintf("Content-Length: %v", len(body)),
		)

		conn.Write(response)
	} else if echoRegExp.MatchString(requestpath){
		var body string

		for index, str := range strings.Split(requestpath, "/") {
			if index < 2 {
				continue
			} else if index == 2 {
				body += str
			} else {
				body += "/" + str
			}
		}
		response := http.BuildResponse(
			200, 
			"OK", 
			body, 
			"Content-Type: text/plain", 
			fmt.Sprintf("Content-Length: %v", len(body)),
		)

		conn.Write(response)
	} else {
		body := "Route not found"
		response := http.BuildResponse(
			404, 
			"Not Found", 
			body, 
			"Content-Type: text/plain", 
			fmt.Sprintf("Content-Length: %v", len(body)),
		)

		conn.Write(response)
	}

	defer conn.Close()
}

func main () {
	fmt.Println("HTTP SERVER RUNNING")

	ln, err := net.Listen("tcp", ":4221")
	if err != nil {
		fmt.Printf("Error binding to a port: %v\n", err)
	}
	var directory string

	flag.StringVar(&directory, "directory", "", "Specify the directory")
	flag.Parse()

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Error accepting TCP connection: %v\n", err)
		}
		
		fmt.Printf("New TCP connection accepted")

		go handleConnection(conn, directory)
	}
}
