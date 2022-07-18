package Delivery

import (
	"encoding/json"
	"net/http"

	"github.com/tieldmoon/tieldauth/Repository"
	"github.com/tieldmoon/tieldauth/Usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func SigninHandler(w http.ResponseWriter, r *http.Request, mongodb *mongo.Client) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	t := Repository.TokenRepositoryMongo{
		Client: mongodb,
	}
	data, available := t.CheckAppIdIsAvailable(r.PostFormValue("app_id"))
	// if not acailable set error
	if !available {
		e, _ := json.Marshal(map[string]any{
			"errorCode": http.StatusNotFound,
			"message":   "Not Found",
		})
		http.Error(w, string(e), http.StatusNotFound)
		return
	}
	// fmt.Println(data, available)
	// if available parsing jwt token
	j, err := Usecase.ParseJWT(r.PostFormValue("secret_key"), data.AppKey)
	if err != nil {
		e, _ := json.Marshal(map[string]any{
			"errorCode": http.StatusBadRequest,
			"message":   err.Error(),
		})
		http.Error(w, string(e), http.StatusBadRequest)
		return
	}
	email := j["email"]
	password := j["password"]
	if email == nil || password == nil {
		e, _ := json.Marshal(map[string]any{
			"errorCode": http.StatusBadRequest,
			"message":   "Invalid jwt payload format",
		})
		http.Error(w, string(e), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Ok"))
}
