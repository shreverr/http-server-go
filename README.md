# Go HTTP Server
This Go HTTP server project provides a lightweight and customizable solution for handling HTTP requests and serving responses. Built using the standard Go library, it offers a simple yet powerful way to create an HTTP server with support for multiple routes and customizable responses.

This is a simple HTTP server implemented in Go.

## Features

- Handles HTTP requests and returns responses.
- Supports multiple routes with customizable responses.
- Easy to configure and extend.

## Installation

To use this HTTP server, you need to have Go installed on your machine.

1. Clone the repository:

   ```bash
   git clone https://github.com/shreverr/http-server-go.git
   cd http-server-go
2. Run the application
   ```bash
   go run main.go

## Usage
Once the server is running, you can access it at http://localhost:4221.

### Routes
- /: Returns a "Hello, World!" message.
- /echo/{message}: Echoes back the provided message.
