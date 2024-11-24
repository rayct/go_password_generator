package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()_+=-{}[]|:;<>,.?/~"
)

var (
	lineCounter int        // Global counter for sequential line numbers
	counterLock sync.Mutex // Mutex to prevent race conditions
)

// Function to log generated passwords
func logPassword(password string) {
	// Open the log file in append mode, create it if it doesn't exist
	logFile, err := os.OpenFile("passwords.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer logFile.Close()

	// Increment the global counter safely
	counterLock.Lock()
	lineCounter++
	currentLine := lineCounter
	counterLock.Unlock()

	// Set up the logger to write to the file
	logger := log.New(logFile, "", log.LstdFlags)

	// Log the password with sequential line number, date, and time
	logger.Printf("Line %d: Generated Password: %s", currentLine, password)
}

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

	// Generate and log passwords multiple times for demonstration
	for i := 0; i < 5; i++ {
		password := generatePassword(passwordLength)
		fmt.Printf("Generated Password: %s\n", password)

		// Log the generated password
		logPassword(password)
	}
}
