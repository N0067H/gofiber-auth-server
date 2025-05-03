package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(h), err
}

func CheckPasswordHash(hashed string, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass))
}
