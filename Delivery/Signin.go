package Delivery

import (
	"net/http"

	"github.com/tieldmoon/tieldauth/Service"
)

func SigninHandler(w http.ResponseWriter, r *http.Request, wo *Service.Worker) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	w.Write([]byte("Ok"))
}
