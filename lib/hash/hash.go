package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(cleartext, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(cleartext))
	return err == nil
}
