package hash

import (
	"golang.org/x/crypto/bcrypt"
)

const HASH_COST = 10

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), HASH_COST)
	return string(bytes), err
}

func CheckPasswordHash(cleartext, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(cleartext))
	return err == nil
}
