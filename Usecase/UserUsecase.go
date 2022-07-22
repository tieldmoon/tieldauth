package Usecase

import (
	"log"
	"github.com/tieldmoon/tieldauth/Models"
	"github.com/tieldmoon/tieldauth/Repository"
	"golang.org/x/crypto/bcrypt"
)

func Login(u Repository.UserRepository, email string, password string) Models.User {
	x, _ := u.FindByEmail(email)
	log.Println(x)
	pw := x.Password
	err := bcrypt.CompareHashAndPassword([]byte(pw), []byte(password))
	if err != nil {
		return Models.User{}
	}
	return x
}
