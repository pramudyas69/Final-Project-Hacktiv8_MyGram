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
func ComparePassword(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
