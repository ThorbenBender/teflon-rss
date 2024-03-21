package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/thorbenbender/teflon-rss/internal/database"
)

func (cfg *apiConfig) HandleFeedFollowCreate(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}
	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	respondWithJson(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))
}

func (cfg *apiConfig) HandleFeedFollowDelete(w http.ResponseWriter, r *http.Request) {
	feedFollowIDString := chi.URLParam(r, "feedFollowID")
	fmt.Println(feedFollowIDString)
	if feedFollowIDString == "" {
		respondWithError(w, http.StatusBadRequest, "No feed follow id found")
		return
	}
	feedFollowID, err := uuid.Parse(feedFollowIDString)
	fmt.Println(feedFollowID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Wrong format")
		return
	}
	err = cfg.DB.DeleteFeedFollow(r.Context(), feedFollowID)
	fmt.Println(err.Error())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}
	respondWithJson(w, http.StatusOK, struct{}{})
}

func (cfg *apiConfig) HandleFeedFollowRetrieve(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	databaseFeedFollows, err := cfg.DB.GetFeedFollowByUserID(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting feed follows")
		return
	}
	feedFollows := make([]FeedFollow, 0, len(databaseFeedFollows))
	for _, feedFollow := range databaseFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(feedFollow))
	}
	respondWithJson(w, http.StatusOK, feedFollows)
}
