package utils

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
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	errResp, _ := json.Marshal(ResponseError{Error: message})
	w.Write(errResp)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	jsonResp, _ := json.Marshal(payload)
	w.Write(jsonResp)
}

func RespondWithPlainText(w http.ResponseWriter, code int, message string) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)
	w.Write([]byte(message))
}

//

func RemoveProfanity(chirpBody string) string {
	badWords := []string{"kerfuffle", "sharbert", "fornax"}

	res := []string{}

	for _, word := range strings.Fields(chirpBody) {
		for _, badWord := range badWords {
			if strings.ToLower(word) == strings.ToLower(badWord) {
				word = "****"
				break
			}
		}
		res = append(res, word)
	}

	return strings.Join(res, " ")
}
