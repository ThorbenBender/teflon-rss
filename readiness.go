package main

import "net/http"

func HandleReadiness(w http.ResponseWriter, r *http.Request) {
	type message struct {
		Status string `json:"status"`
	}

	msg := message{
		Status: "ok",
	}
	respondWithJson(w, http.StatusOK, msg)
}
