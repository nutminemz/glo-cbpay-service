package utility

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPin(pin string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pin), 6)
	return string(bytes), err
}

func CheckPinHash(pin string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pin))
	return err == nil
}
