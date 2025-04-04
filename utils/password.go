package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hashpasss, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(hashpasss), err
}

func ComparePass(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
