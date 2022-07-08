package Delivery

import (
	"fmt"
	"net/http"

	"github.com/tieldmoon/tieldauth/Repository"
	"go.mongodb.org/mongo-driver/mongo"
)

func SigninHandler(w http.ResponseWriter, r *http.Request, mongodb *mongo.Client) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	t := Repository.TokenRepositoryMongo{
		Client: mongodb,
	}
	available := t.CheckAppIdIsAvailable(r.PostFormValue("app_id"))
	fmt.Println(available)

	// m := <-wo.Mongo
	// m.Ping(context.TODO(), readpref.Primary())
	w.Write([]byte("Ok"))
}
