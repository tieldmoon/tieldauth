package Usecase

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func ParseJWT(secret string, appkey string) (jwt.MapClaims, error) {
	t, err := verifyJWT(secret, appkey)
	if err != nil {
		return nil, err
	}
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid Jwt token")
}

func verifyJWT(secret string, appkey string) (*jwt.Token, error) {
	token, err := jwt.Parse(secret, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != jwt.SigningMethodHS512 {
			return nil, fmt.Errorf("signing method invalid")
		}
		return []byte(appkey), nil
	})
	return token, err
}
