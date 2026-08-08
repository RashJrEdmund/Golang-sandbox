/*
SEE WRITING struct of a request handler here
	https://pkg.go.dev/net/http#ResponseWriter.WriteHeader
*/

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/RashJrEdmund/go-sandbox/chirpy/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

// MIDDLEWARES

type apiConfig struct {
	// The atomic.Int32 type is a really cool standard-library type that allows us to safely increment and read
	// an integer value across multiple goroutines (HTTP requests). https://pkg.go.dev/sync/atomic#Int32
	fileServerHits atomic.Int32
	dbQueries      *database.Queries
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
	w.Header().Add("Content-Type", "text/html; charset=utf-8")

	restTemplate := fmt.Sprintf(`
	<html>
		<body>
			<h1>Welcome, Chirpy Admin</h1>
			<p>Chirpy has been visited %d times!</p>
		</body>
	</html>
	`, apiCfg.fileServerHits.Load())
	w.Write([]byte(restTemplate))
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

func validateChirpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	var data Chirp

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&data); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Something went wrong")
		return
	}

	if len(data.Body) > 140 {
		RespondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}

	RemoveProfanity(&data)

	RespondWithJSON(w, http.StatusOK, ValidateChirpResponse{CleanedBody: data.Body})
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	mu := http.NewServeMux()

	const PORT = "8080"

	const rootDir = "."

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: mu,
	}

	apiCfg := &apiConfig{
		fileServerHits: atomic.Int32{},
		dbQueries:      database.New(db),
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

	mu.Handle("GET /admin/metrics/", http.HandlerFunc(apiCfg.metricsHandler))

	mu.Handle("POST /admin/reset/", http.HandlerFunc(apiCfg.resetHandler))

	mu.HandleFunc("GET /api/healthz/", healthzHandler)

	mu.HandleFunc("POST /api/validate_chirp/", validateChirpHandler)

	fmt.Printf("Serving files from %s on port %s\n", rootDir, PORT)
	log.Fatal(server.ListenAndServe())
}
