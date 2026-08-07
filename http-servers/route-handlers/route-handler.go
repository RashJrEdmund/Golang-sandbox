package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

/*
DIFFERENT WAYS TO REGISTER A HANDLER WITH A SERVER
	Read: https://pkg.go.dev/net/http#ServeMux
*/

type data struct {
	Message string `json:"message"`
	Route   string `json:"route"`
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(4)
	go firstWay(wg)
	go secondWay(wg)
	go thirdWay(wg)
	go nestedRouters(wg)
	wg.Wait()
}

// --------------------------------------------------------------
// FIRST WAY: USING http.Handler with a struct.
// --------------------------------------------------------------

type userHandler struct{} // implements the http.Handler interface by having a ServeHTTP method

func (userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := data{
		Message: "Hello, World. This is the first way to register a handler with a server.",
		Route:   r.URL.Path,
	}

	res, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	fmt.Println("First way: Handler executed")
}

func firstWay(wg *sync.WaitGroup) { // using the serveMux.Handle with a struct method
	defer wg.Done()

	serveMux := http.NewServeMux()

	server := http.Server{
		Addr:    ":8080", // If empty, ":http" (port 80) is used.
		Handler: serveMux,
	}

	serveMux.Handle("/first-way", userHandler{})

	fmt.Println("First way: Server is running on port 8080")
	fmt.Println(Delimiter)
	server.ListenAndServe() // this blocks the main thread until the server is stopped
	fmt.Println("Not reached")
}

// --------------------------------------------------------------
// SECOND WAY: USING http.HandlerFunc with a function.
// --------------------------------------------------------------

func secondWayHandler(w http.ResponseWriter, r *http.Request) {
	data := data{
		Message: "Hello, World. This is the second way to register a handler with a server.",
		Route:   r.URL.Path,
	}

	res, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	fmt.Println("Second way: Handler executed")
}

func secondWay(wg *sync.WaitGroup) {
	//  http.HandlerFunc is a function that takes a http.ResponseWriter and a http.Request and returns a http.Handler.j
	// so no need to create struct that implements the http.Handler interface.
	defer wg.Done()

	serveMux := http.NewServeMux()

	server := http.Server{
		Addr:    ":8081", // If empty, ":http" (port 80) is used.
		Handler: serveMux,
	}

	serveMux.Handle("/second-way", http.HandlerFunc(secondWayHandler))

	fmt.Println("Second way: Server is running on port 8081")
	fmt.Println(Delimiter)
	server.ListenAndServe()
}

/*
Under the hood, http.HandlerFunc does

	type HandlerFunc func(ResponseWriter, *Request)

	//Then it gives that type a method:

	func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
		f(w, r)
	}

	essentially utilizing ServeHTTP method
*/

// --------------------------------------------------------------
// THIRD WAY: USING http.HandleFunc directly with a function.
// --------------------------------------------------------------

func thirdWayHandler(w http.ResponseWriter, r *http.Request) {
	data := data{
		Message: "Hello, World. This is the third way to register a handler with a server.",
		Route:   r.URL.Path,
	}

	res, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	fmt.Println("Third way: Handler executed")
}

func thirdWay(wg *sync.WaitGroup) {
	defer wg.Done()

	serveMux := http.NewServeMux()

	server := http.Server{
		Addr:    ":8082", // If empty, ":http" (port 80) is used.
		Handler: serveMux,
	}

	serveMux.HandleFunc("/third-way", thirdWayHandler)

	fmt.Println("Third way: Server is running on port 8082")
	fmt.Println(Delimiter)
	server.ListenAndServe()
}

/*
Under the hood, http.HandleFunc does

something like: mux.Handle("/hello", http.HandlerFunc(hello))
*/

// --------------------------------------------------------------
// NESTED ROUTERS
// --------------------------------------------------------------

/*
	Here's something cool

// Because *http.ServeMux implements http.Handler, you can nest routers:
The admin mux is passed directly to Handle because it satisfies the http.Handler interface.
*/

func usersHandler(w http.ResponseWriter, _ *http.Request) {
	data := data{
		Message: "Hello, World. This is the admin users page.",
		Route:   "/users",
	}

	res, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	fmt.Println("Nested routers: Users handler executed")
}

func nestedRouters(wg *sync.WaitGroup) {
	defer wg.Done()

	adminMu := http.NewServeMux()
	adminMu.HandleFunc("/users", usersHandler)

	rootMu := http.NewServeMux()
	rootMu.Handle("/admin/", adminMu)

	server := http.Server{
		Addr:    ":8083", // If empty, ":http" (port 80) is used.
		Handler: rootMu,
	}

	fmt.Println("Nested routers: Server is running on port 8083")
	fmt.Println(Delimiter)
	server.ListenAndServe()
}
