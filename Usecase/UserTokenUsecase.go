package Usecase

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/tieldmoon/tieldauth/Models"
)

func GenerateUserToken(secretkey string, user Models.User) (string, error) {
	claims := Models.UserTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    user.Username,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Duration(2) * time.Hour).Unix(),
		},
		Username: user.Username,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signed, err := t.SignedString([]byte(secretkey))
	if err != nil {
		return "", err
	}
	return signed, err
}

func GenerateRefreshToken(secretkey string, user Models.User) (string, error) {
	claims := jwt.StandardClaims{
		Issuer:    user.Username,
		ExpiresAt: time.Now().Add(time.Duration(336) * time.Hour).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signed, err := t.SignedString([]byte(secretkey))
	if err != nil {
		return "", err
	}
	return signed, err
}
