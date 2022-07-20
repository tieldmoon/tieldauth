package Usecase

import (
	"github.com/tieldmoon/tieldauth/Repository"
	"golang.org/x/crypto/bcrypt"
)

func Login(u Repository.UserRepository, email string, password string) bool {
	x, _ := u.FindByEmail(email)
	pw := x.Password
	err := bcrypt.CompareHashAndPassword([]byte(pw), []byte(password))
	if err != nil {
		return false
	}
	return true
}
