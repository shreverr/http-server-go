package http

import (
	"strings"
)

func GetRequestPath (request string) string {
	requestPath := strings.Split(strings.Split(request, "\r\n")[0], " ")[1]
	
	return requestPath
}