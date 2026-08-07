/*
For example, we can write a middleware that logs every request to the server.
We can then wrap our handler with this middleware and every request will be logged without
having to write the logging code in every handler.
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request
		fmt.Println(Delimiter)
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
