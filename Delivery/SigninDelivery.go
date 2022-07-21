package Delivery

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/tieldmoon/tieldauth/Repository"
	"github.com/tieldmoon/tieldauth/Usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

// Signin Handler
//
// /api/oauth2/signin
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

	password := j["password"].(string)
	u := Repository.UserRepositoryMongo{
		Client: mongodb,
	}
	// verify email password
	user := Usecase.Login(&u, email.(string), string(password))

	// user token
	usertoken, err := Usecase.GenerateUserToken(data.AppKey, user)
	if err != nil {
		e, _ := json.Marshal(map[string]any{
			"statusCode":    http.StatusConflict,
			"message":       "error generate usertoken",
			"user_token":    "",
			"refresh_token": "",
		})
		w.Write([]byte(e))
		return
	}

	// refresh token
	refreshtoken, err := Usecase.GenerateRefreshToken(data.AppKey, user)
	if err != nil {
		e, _ := json.Marshal(map[string]any{
			"statusCode":    http.StatusConflict,
			"message":       "error generate refreshtoken",
			"user_token":    "",
			"refresh_token": "",
		})
		w.Write([]byte(e))
		return
	}
	if !reflect.ValueOf(user).IsZero() {
		// generate user token and refresh token
		e, _ := json.Marshal(map[string]any{
			"statusCode":    http.StatusOK,
			"message":       "login success",
			"user_token":    usertoken,
			"refresh_token": refreshtoken,
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
