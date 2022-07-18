package Usecase

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
)

func ParseJWT(secret string, appkey string) error {
	t, err := verifyJWT(secret, appkey)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(t)
	return nil
}

func verifyJWT(secret string, appkey string) (*jwt.Token, error) {
	token, err := jwt.Parse(secret, func(t *jwt.Token) (interface{}, error) {
		if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != jwt.SigningMethodHS512 {
			return nil, fmt.Errorf("signing method invalid")
		}
		return []byte("eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTZzNzg5MCIsIm5hbWUiOiJKb2hzbiBEb2UiLCJpYXQiOjE1MTYyMzkwMjJ9.beQwR-o21Fh2VYaqlE8hKQkjryZrU4IruFmHthLFKKW3uA4Bl0MK3sru0B_1wj2eBNGw9h5DMISJBygd7Jnulg"), nil
	})
	return token, err
}
