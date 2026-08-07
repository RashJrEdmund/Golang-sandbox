/*
File server is a simple HTTP server that serves files from a directory.
It's contents are rooted from the root directory of the server.
	Read more: https://pkg.go.dev/net/http#FileServer

Normally to see a file, you'll have to write it's path filename and .extension.
	But for index.html, go handles it automatically and no need for filename
*/

package main

import (
	"fmt"
	"net/http"
)

// const rootDir = "/home/rash/Desktop/web-dev/projects/go-sandbox/http-servers/file-server/"
const rootDir = "." // current directory

func main() {
	mu := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mu,
	}

	mu.Handle("/", http.FileServer(http.Dir(rootDir)))

	fmt.Println("Server is running on port 8080")
	server.ListenAndServe()
}
