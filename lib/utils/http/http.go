package http

import (
	"strings"
	"fmt"
)

func GetRequestPath (request string) string {
	requestPath := strings.Split(strings.Split(request, "\r\n")[0], " ")[1]
	
	return requestPath
}

func BuildResponse (statusCode int, statusText string, body string, headers ...string) []byte {
	var headerString string
	for _, header := range headers {
		headerString += fmt.Sprintf("%v\r\n", header) 
	}
	
	formattedResponse := fmt.Sprintf(
		"HTTP/1.1 %v %v\r\n%v\r\n%v",
		statusCode,
		statusText,
		headerString,
		body)

		fmt.Printf("Formatted Response: %q", formattedResponse)
	return []byte(formattedResponse)
}
