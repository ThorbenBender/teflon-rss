package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/thorbenbender/teflon-rss/internal/database"
)

func (cfg *apiConfig) HandlePostRetrieve(
	w http.ResponseWriter,
	r *http.Request,
	user database.User,
) {
	var limit int32 = 10
	query := r.URL.Query().Get("limit")
	if query != "" {
		limitConv, err := strconv.Atoi(query)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid limit")
			return
		}
		log.Printf("Conversed limit is %d", limitConv)
		limit = int32(limitConv)
	}
	log.Printf("Limit is %d", limit)
	databasePosts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couln't get posts")
		return
	}

	respondWithJson(w, http.StatusOK, databasePostsToPosts(databasePosts))
}
