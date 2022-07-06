package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/tieldmoon/tieldauth/Delivery"
	"github.com/tieldmoon/tieldauth/Service"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	godotenv.Load()

	Workers := Service.Worker{
		Wg:    new(sync.WaitGroup),
		Jobs:  make(chan map[int]interface{}, 100),
		Mongo: make(chan *mongo.Client, 100),
	}
	go Workers.InitWorker(100)
	Workers.Wg.Wait()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Oauth2
	r.Group(func(r chi.Router) {
		r.Post("/api/oauth2/signin", func(w http.ResponseWriter, r *http.Request) {
			Delivery.SigninHandler(w, r, &Workers)
		})
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.ListenAndServe(":"+port, r)
}
