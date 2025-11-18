package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(passwordString string) string {
	passwordBytes := []byte(passwordString)

	// Generate bcrypt hash
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}

func MatchesHash(passwordString string, hashString string) bool {
	passwordBytes := []byte(passwordString)
	hashBytes := []byte(hashString)

	err := bcrypt.CompareHashAndPassword(hashBytes, passwordBytes)
	if err != nil {
		fmt.Println("Password does not match")
		return false
	}
	fmt.Println("Password matches")
	return true
}
