package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

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
		fmt.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldnt create user")
		return
	}
	respondWithJson(w, http.StatusCreated, databaseUserToUser(user))
}

func (cfg *apiConfig) HandleUserRetrieve(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	respondWithJson(w, http.StatusOK, databaseUserToUser(user))
}
