package Delivery

import (
	"net/http"
)

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	w.Write([]byte("Ok"))
}
