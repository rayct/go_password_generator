## Password Generator written in Go - Release v0.1.0

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [About](#about)


### Overview

The Password Generator is a command-line program written in Go (Golang) that creates strong, random passwords of variable length. It uses a combination of uppercase letters, lowercase letters, digits, and special characters to generate passwords.

### Installation

1. Ensure you have Go installed on your machine.
2. Copy the provided Go code into a file named `password_generator.go`.
3. Run the program using `go run password_generator.go`.

### Usage

#### Generating Passwords

The program allows you to specify the length of the password to generate. Follow these steps to generate a password:

1. Locate the `passwordLength` variable in the code.
2. Modify the value assigned to `passwordLength` to set the desired length of the password.
3. Run the program. It will generate a password of the specified length and display it in the console.

### Customization

#### Changing Password Length

You can adjust the length of the generated password by modifying the `passwordLength` variable in the `main()` function. Simply assign a different integer value to `passwordLength` to change the length of the generated password.

### Security Considerations

The generated passwords are strong and random due to the use of various character types (uppercase letters, lowercase letters, digits, and special characters). However, ensure that generated passwords are stored securely and not transmitted or stored in plaintext.

### Code Structure

The program consists of the following key components:

- **generatePassword() Function**: Generates a random password based on the specified length and character pool.
- **Constants**: Defines character sets for uppercase letters, lowercase letters, digits, and special characters.
- **Main Function**: Controls the execution of the program. It sets the desired password length and displays the generated password in the console.

### Example

```go
package main

// ... (Code provided earlier)

func main() {
    passwordLength := 18 // Change this value to set the length of the password
    password := generatePassword(passwordLength)
    fmt.Printf("Generated Password: %s\n", password)
}
```

### Note

- For enhanced security, consider implementing additional measures such as securely storing generated passwords and avoiding exposure in unencrypted forms.

---

Documentation By: **Raymond C. TURNER**

**Revision:** November 24th, 2023

**rayturner.dev**