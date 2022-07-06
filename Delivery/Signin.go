package Delivery

import (
	"net/http"

	"github.com/tieldmoon/tieldauth/Service"
)

func SigninHandler(w http.ResponseWriter, r *http.Request, wo *Service.Worker) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	wo.Jobs <- map[int]any{1: "hello"}
	w.Write([]byte("Ok"))
}
