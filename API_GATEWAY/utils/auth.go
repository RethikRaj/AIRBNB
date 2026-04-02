package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CompareAndVerifyPassword(password_hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password))

	return err == nil // err == nil => true
}
