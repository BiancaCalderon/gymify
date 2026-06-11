package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	"github.com/biancacalderon/gymify-backend/internal/config"
	"github.com/biancacalderon/gymify-backend/internal/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró archivo .env, usando variables de entorno del sistema")
	}

	cfg := config.Load()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	// Health check
	r.Get("/health", handlers.HealthCheck)

	// API v1
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", handlers.Register)
			r.Post("/login", handlers.Login)
		})

		r.Route("/users", func(r chi.Router) {
			r.Get("/{id}", handlers.GetUser)
		})

		r.Route("/activities", func(r chi.Router) {
			r.Post("/", handlers.CreateActivity)
			r.Get("/user/{id}", handlers.GetUserActivities)
		})

		r.Route("/streaks", func(r chi.Router) {
			r.Get("/user/{id}", handlers.GetUserStreak)
		})

		r.Route("/feed", func(r chi.Router) {
			r.Get("/", handlers.GetFeed)
			r.Post("/posts", handlers.CreatePost)
		})

		r.Route("/challenges", func(r chi.Router) {
			r.Get("/", handlers.GetChallenges)
			r.Post("/{id}/join", handlers.JoinChallenge)
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Gymify backend corriendo en el puerto %s (env: %s)", port, cfg.Env)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
