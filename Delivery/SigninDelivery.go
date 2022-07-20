package Delivery

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tieldmoon/tieldauth/Repository"
	"github.com/tieldmoon/tieldauth/Usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func SigninHandler(w http.ResponseWriter, r *http.Request, mongodb *mongo.Client) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	t := Repository.TokenRepositoryMongo{
		Client: mongodb,
	}
	data, available := t.FindAppId(r.PostFormValue("app_id"))
	// if not acailable set error
	if !available {
		e, _ := json.Marshal(map[string]any{
			"statusCode": http.StatusNotFound,
			"message":    "Not Found",
		})
		http.Error(w, string(e), http.StatusNotFound)
		return
	}
	// fmt.Println(data, available)
	// if available parsing jwt token
	j, err := Usecase.ParseJWT(r.PostFormValue("secret_key"), data.AppKey)
	if err != nil {
		e, _ := json.Marshal(map[string]any{
			"statusCode": http.StatusBadRequest,
			"message":    err.Error(),
		})
		http.Error(w, string(e), http.StatusBadRequest)
		return
	}
	email := j["email"]
	if email == nil || j["password"] == nil {
		e, _ := json.Marshal(map[string]any{
			"statusCode": http.StatusBadRequest,
			"message":    "Invalid jwt payload format",
		})
		http.Error(w, string(e), http.StatusBadRequest)
		return
	}
	a, _ := bcrypt.GenerateFromPassword([]byte(j["password"].(string)), 13)
	log.Println(string(a))
	password := j["password"].(string)
	u := Repository.UserRepositoryMongo{
		Client: mongodb,
	}
	// verify email password
	if success := Usecase.Login(&u, email.(string), string(password)); success {
		// generate user token and refresh token
		e, _ := json.Marshal(map[string]any{
			"statusCode":    http.StatusOK,
			"message":       "login success",
			"user_token":    "",
			"refresh_token": "",
		})
		w.Write([]byte(e))
		return
	}

	// result
	e, _ := json.Marshal(map[string]any{
		"statusCode":    http.StatusFound,
		"message":       "invalid email or password",
		"user_token":    "",
		"refresh_token": "",
	})
	w.Write([]byte(e))
}
