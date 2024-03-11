package main

import "net/http"

func HandleError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Secret message")
}
