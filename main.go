package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

var nowFunc = time.Now

const (
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()_+=-{}[]|:;<>,.?/~"
)

var (
	lineCounter int
	counterLock sync.Mutex

	// logPassword is a function variable for easy override/swapping
	logPassword func(string)
)

func init() {
	// Default logPassword uses UK format
	setLogFormat("UK")
}

func setLogFormat(format string) {
	switch strings.ToUpper(format) {
	case "US":
		logPassword = func(password string) {
			logFile, err := os.OpenFile("passwords.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("Error opening log file: %v\n", err)
				return
			}
			defer logFile.Close()

			counterLock.Lock()
			lineCounter++
			currentLine := lineCounter
			counterLock.Unlock()

			currentTime := nowFunc().Format("01/02/2006 03:04:05 PM") // MM/DD/YYYY hh:mm:ss AM/PM

			logger := log.New(logFile, "", 0)
			logger.Printf("%s - Line %d: Generated Password: %s", currentTime, currentLine, password)
		}
	case "UK":
		fallthrough
	default:
		logPassword = func(password string) {
			logFile, err := os.OpenFile("passwords.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("Error opening log file: %v\n", err)
				return
			}
			defer logFile.Close()

			counterLock.Lock()
			lineCounter++
			currentLine := lineCounter
			counterLock.Unlock()

			currentTime := time.Now().Format("02/01/2006 15:04:05") // DD/MM/YYYY 24-hour

			logger := log.New(logFile, "", 0)
			logger.Printf("%s - Line %d: Generated Password: %s", currentTime, currentLine, password)
		}
	}
}

func generatePassword(length int) string {
	characterPool := uppercaseLetters + lowercaseLetters + digits + specialChars
	rand.Seed(time.Now().UnixNano())

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(characterPool))
		password[i] = characterPool[randomIndex]
	}
	return string(password)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose date format (US/UK) [default UK]: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input error, defaulting to UK format")
		input = "UK"
	}
	input = strings.TrimSpace(input)
	if input == "" {
		input = "UK"
	}

	setLogFormat(input)

	passwordLength := 22
	for i := 0; i < 5; i++ {
		password := generatePassword(passwordLength)
		fmt.Printf("Generated Password: %s\n", password)
		logPassword(password)
	}
}
