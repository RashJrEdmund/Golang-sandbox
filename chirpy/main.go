/*
SEE WRITING struct of a request handler here
	https://pkg.go.dev/net/http#ResponseWriter.WriteHeader
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

// MIDDLEWARES

type apiConfig struct {
	// The atomic.Int32 type is a really cool standard-library type that allows us to safely increment and read
	// an integer value across multiple goroutines (HTTP requests). https://pkg.go.dev/sync/atomic#Int32
	fileServerHits atomic.Int32
}

func (apiCfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Incrementing metrics")
		apiCfg.fileServerHits.Add(1)

		fmt.Println(apiCfg.fileServerHits.Load())
		next.ServeHTTP(w, r)
	})
}

func (apiCfg *apiConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")

	restText := fmt.Sprintf("Hits: %d", apiCfg.fileServerHits.Load())
	w.Write([]byte(restText))
}

func (apiCfg *apiConfig) resetHandler(w http.ResponseWriter, r *http.Request) {
	apiCfg.fileServerHits.Store(0)

	resText := fmt.Sprintf("Hits: %d", apiCfg.fileServerHits.Load())

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(resText))
}

// ROUTE HANDLERS

type rootHandler struct{}

func (rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

// ------------------Health Handler----------------------------------------

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("OK")) // w.Write([]byte(http.StatusText(http.StatusOK)))
}

func main() {
	mu := http.NewServeMux()

	const PORT = "8080"

	const rootDir = "."

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: mu,
	}

	apiCfg := &apiConfig{
		fileServerHits: atomic.Int32{},
	}

	/*
		Now that the path is no longer "/", we need to fix this by using http.StripPrefix
			Read here: https://pkg.go.dev/net/http#StripPrefix
	*/
	mu.Handle("/app/",
		apiCfg.middlewareMetricsInc(
			http.StripPrefix("/app/", http.FileServer(http.Dir(rootDir))),
		),
	)

	mu.Handle("/metrics/", http.HandlerFunc(apiCfg.metricsHandler))

	mu.Handle("/reset/", http.HandlerFunc(apiCfg.resetHandler))

	mu.HandleFunc("/healthz/", healthzHandler)

	fmt.Printf("Serving files from %s on port %s\n", rootDir, PORT)
	log.Fatal(server.ListenAndServe())
}
