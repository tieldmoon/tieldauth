package main

import (
	"net/http"
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/tieldmoon/tieldauth/Delivery"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	err := godotenv.Load()
	port := "3000"
	if err == nil {
		port = os.Getenv("PORT")
	}
	worker := Worker{
		Wg:    new(sync.WaitGroup),
		Jobs:  make(chan map[int]interface{}, 100),
		Mongo: make(chan *mongo.Client, 100),
	}
	go worker.InitWorker(100)
	worker.Wg.Wait()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Oauth2
	r.Group(func(r chi.Router) {
		r.Post("/api/oauth2/signin", Delivery.SigninHandler)
	})
	http.ListenAndServe(":"+port, r)
}
