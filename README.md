## CLI Password Generator written in Go - Release v0.2.0

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Logging Feature](#logging-feature)
- [Customization](#customization)
- [Security Considerations](#security-considerations)
- [Code Structure](#code-structure)
- [Example](#example)
- [Contributing](#contributing)
- [License](#license)
- [About](#about)

---

### Introduction

The Password Generator is a command-line program written in Go (Golang) that creates strong, random passwords of variable length. It uses a combination of uppercase letters, lowercase letters, digits, and special characters to generate passwords. As of version **v0.1.0**, it includes a logging feature to store generated passwords with timestamps and sequential line numbers.

---

### Features

- Generates strong, random passwords.
- Supports customizable password lengths.
- Logs all generated passwords to a file (`passwords.log`) with:
  - Sequential line numbers.
  - Date and time of generation.

---

### Requirements

- Go 1.18+ installed on your machine.

---

### Installation

1. Ensure you have Go installed.
2. Copy the provided Go code into a file named `password_generator.go`.
3. Run the program using `go run password_generator.go`.

---

### Usage

#### Generating Passwords

1. Open `password_generator.go`.
2. Set the desired password length by modifying the `passwordLength` variable.
3. Run the program: `go run password_generator.go`.
4. View the generated password in the console.

#### Logging Feature

All generated passwords are automatically logged to a file named `passwords.log` in the same directory as the program. Each log entry includes:
- A sequential line number.
- The generated password.
- The date and time of generation.

The logging system ensures traceability of generated passwords:
- **File**: `passwords.log`.
- **Format**:
  ```
  YYYY/MM/DD HH:MM:SS Line <Line Number>: Generated Password: <Password>
  ```

### Feature Changes

#### v0.1.0 - First inital Release
- Logging Passwords to txt file, Date and time stamp with sequential line numbers.
#### v0.2.0
- Changed Date format to from US to UK.


#### Example Log Entry:
```
2024/11/24 15:55:32 Line 1: Generated Password: Y@3pL0w*Xq7Za!R5C9
2024/11/24 15:55:33 Line 2: Generated Password: Wt9@Xp2$#Rq6Lb8Y!K
```

---

### Customization

#### Changing Password Length
Modify the `passwordLength` variable in the `main()` function to set your desired password length.

#### Disabling Logging
To disable logging, comment out or remove the `logPassword` call in the `main()` function.

---

### Security Considerations

- The program generates strong and random passwords using a mix of uppercase letters, lowercase letters, digits, and special characters.
- Generated passwords are logged in plaintext in `passwords.log`. Ensure the file is stored securely and access is restricted.

---

### Code Structure

The program consists of the following key components:

- **`generatePassword()` Function**: Generates a random password based on the specified length and character pool.
- **`logPassword()` Function**: Logs generated passwords to `passwords.log` with sequential line numbers and timestamps.
- **Constants**: Define character sets for uppercase letters, lowercase letters, digits, and special characters.
- **Main Function**: Controls program execution, including setting password length and triggering logging.

---

### Example

```go
package main

// ... (Code provided earlier)

func main() {
    passwordLength := 22 // Set the desired password length
    password := generatePassword(passwordLength)
    fmt.Printf("Generated Password: %s\n", password)

    // Logs the password to passwords.log
    logPassword(password)
}
```

---

### Contributing

Contributions are welcome! Submit issues or pull requests on the [GitHub repository](#).

---

### License

MIT License. See `LICENSE` file for details.

---

### About

Documentation By: **Raymond C. Turner**

**Revision:** August 8th, 2025

**Author:** [rayturner.dev]()