package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	// Set the password length for testing
	passwordLength := 10

	// Generate password
	password := generatePassword(passwordLength)

	// Log generated password to a text file
	fileName := "generated_passwords.txt"
	err := logPasswordToFile(fileName, password)
	if err != nil {
		t.Errorf("Error logging password to file: %s", err)
	}

	// Read the file and check if the password exists in the file
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Errorf("Error reading file: %s", err)
	}

	if !strings.Contains(string(fileContent), password) {
		t.Errorf("Generated password not found in the file")
	}
}

func logPasswordToFile(fileName, password string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Generated Password: %s\n", password)
	if err != nil {
		return err
	}

	return nil
}
