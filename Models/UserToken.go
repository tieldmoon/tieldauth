package Models

import "github.com/golang-jwt/jwt"

type UserToken struct {
	User_token    string `bson:"user_token"`
	Refresh_token string `bson:"refresh_token"`
	Expired       string `bson:"expired"`
}

type UserTokenClaims struct {
	jwt.StandardClaims
	Username string `json:"Username"`
}
