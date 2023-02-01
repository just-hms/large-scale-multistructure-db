package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// given a password it hashes and salt it
func HashAndSalt(password string) (string, error) {

	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hash), err
}

// given an hash and a plain password
// returns true if they match
func Verify(hashedPassword string, plainPassword string) bool {

	bytePlain := []byte(plainPassword)
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)

	return err == nil
}
