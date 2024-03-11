package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	data, err := json.Marshal(payload)
	if err != nil {
		w.Write([]byte("Something went wrong"))
	}
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, status int, msg string) {
	if status > 499 {
		log.Printf("Error: %s", msg)
		msg = "Internal Server Error"
	}

	type errorResponse struct {
		Error string `json:"error"`
	}
	errResponse := errorResponse{
		Error: msg,
	}
	respondWithJson(w, status, errResponse)
}
