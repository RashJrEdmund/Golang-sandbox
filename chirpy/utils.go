package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ResponseError struct {
	Error string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json;")

	errResp, _ := json.Marshal(ResponseError{Error: message})
	w.Write(errResp)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json;")
	w.WriteHeader(code)

	jsonResp, _ := json.Marshal(payload)
	w.Write(jsonResp)
}

//

func RemoveProfanity(chirp *Chirp) {
	badWords := []string{"kerfuffle", "sharbert", "fornax"}

	res := []string{}

	for _, word := range strings.Fields(chirp.Body) {
		for _, badWord := range badWords {
			if strings.ToLower(word) == strings.ToLower(badWord) {
				word = "****"
				break
			}
		}
		res = append(res, word)
	}

	chirp.Body = strings.Join(res, " ")
}
