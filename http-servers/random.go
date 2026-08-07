/*
The go standard library makes  it easy to implement a simple server.
	-> Read: https://pkg.go.dev/net/http#hdr-Servers

The http (net/http) package already provides an http.Server for us
	-> Read: https://pkg.go.dev/net/http#Server
*/

package main

import (
	"fmt"
	"time"
)

func handleRequests(reqsChan <-chan request) {
	for req := range reqsChan {
		go handleRequest(req)
	}
}

// don't touch below this line

type request struct {
	path string
}

func main() {
	reqs := make(chan request, 100)
	go handleRequests(reqs)
	for i := range 4 {
		reqs <- request{path: fmt.Sprintf("/path/%d", i)}
		time.Sleep(500 * time.Millisecond)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("5 seconds passed, killing server")
}

func handleRequest(req request) {
	fmt.Println("Handling request for", req.path)
	time.Sleep(2 * time.Second)
	fmt.Println("Done with request for", req.path)
}
