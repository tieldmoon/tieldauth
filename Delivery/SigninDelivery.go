package Delivery

import (
	"net/http"

	"github.com/tieldmoon/tieldauth/Repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func SigninHandler(w http.ResponseWriter, r *http.Request, mongodb *mongo.Client) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	_ = Repository.TokenRepositoryMongo{
		Client: mongodb,
	}

	// m := <-wo.Mongo
	// m.Ping(context.TODO(), readpref.Primary())
	w.Write([]byte("Ok"))
}
