package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(passsword string) string {
	pw := []byte(passsword)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(result)
}
func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(password)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}
