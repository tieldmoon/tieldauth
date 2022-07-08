package main

import (
	"context"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/tieldmoon/tieldauth/Service"
)

func main() {
	godotenv.Load()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db := Service.MongoConnPool()
	defer db.Disconnect(context.TODO())

	// Oauth2
	r.Group(func(r chi.Router) {
		r.Post("/api/oauth2/signin", func(w http.ResponseWriter, r *http.Request) {

		})
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, r)
}
