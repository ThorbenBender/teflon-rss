package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/thorbenbender/teflon-rss/internal/auth"
	"github.com/thorbenbender/teflon-rss/internal/database"
)

func (cfg *apiConfig) HandleUserCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "User name missing")
		return
	}
	createUserParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	}
	user, err := cfg.DB.CreateUser(r.Context(), createUserParams)
	if err != nil {
		log.Fatal(err)
		respondWithError(w, http.StatusInternalServerError, "Couldnt create user")
		return
	}
	respondWithJson(w, http.StatusCreated, databaseUserToUser(user))
}

func (cfg *apiConfig) HandleUserGet(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Could not find user")
		return
	}
	respondWithJson(w, http.StatusOK, databaseUserToUser(user))
}
