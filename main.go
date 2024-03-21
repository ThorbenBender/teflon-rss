package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/thorbenbender/teflon-rss/internal/database"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("No DB URL")
	}
	if port == "" {
		log.Fatal("No port given")
	}
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	apiCfg := apiConfig{
		DB: dbQueries,
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Authorization", "Accept", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	subRouter := chi.NewRouter()
	subRouter.Get("/readiness", HandleReadiness)
	subRouter.Get("/err", HandleError)
	subRouter.Post("/users", apiCfg.HandleUserCreate)
	subRouter.Get("/users", apiCfg.middlewareAuth(apiCfg.HandleUserRetrieve))
	subRouter.Post("/feeds", apiCfg.middlewareAuth(apiCfg.HandleFeedCreate))
	subRouter.Get("/feeds", apiCfg.HandleFeedsRetrieve)
	subRouter.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.HandleFeedFollowCreate))
	subRouter.Delete("/feed_follows/{feedFollowID}", apiCfg.HandleFeedFollowDelete)
	subRouter.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.HandleFeedFollowRetrieve))

	router.Mount("/v1", subRouter)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Printf("Server listening on port %s", port)

	log.Fatal(server.ListenAndServe())
}
