package Delivery

import (
	"fmt"
	"net/http"

	"github.com/tieldmoon/tieldauth/Repository"
	"github.com/tieldmoon/tieldauth/Service"
)

func SigninHandler(w http.ResponseWriter, r *http.Request, wo *Service.Worker) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	appid := r.Form.Get("app_id")
	fmt.Println(appid)
	_ = Repository.TokenRepositoryMongo{
		Client: <-wo.Mongo,
	}
	// _ = <-wo.Mongo
	w.Write([]byte("Ok"))
}
