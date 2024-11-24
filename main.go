package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()_+=-{}[]|:;<>,.?/~"
)

func generatePassword(length int) string {
	// Define characters pool for the password
	characterPool := uppercaseLetters + lowercaseLetters + digits + specialChars

	rand.Seed(time.Now().UnixNano())

	// Generate the password
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(characterPool))
		password[i] = characterPool[randomIndex]
	}

	return string(password)
}

func main() {
	passwordLength := 22 // Change this value to set the length of the password
	password := generatePassword(passwordLength)
	fmt.Printf("Generated Password: %s\n", password)
}
