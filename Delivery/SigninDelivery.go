package Delivery

import (
	"fmt"
	"net/http"

	"github.com/tieldmoon/tieldauth/Service"
)

func SigninHandler(w http.ResponseWriter, r *http.Request, wo *Service.Worker) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	wo.Wg.Add(1)
	wo.Jobs <- map[int]any{1: "hello world"}

	appid := r.Form.Get("app_id")
	fmt.Println(appid)

	// m := <-wo.Mongo
	// m.Ping(context.TODO(), readpref.Primary())
	w.Write([]byte("Ok"))
}