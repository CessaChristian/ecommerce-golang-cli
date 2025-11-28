package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/term"
)

// ReadInput: input string biasa
func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ReadPassword() string {
    fd := int(os.Stdin.Fd())

    // Nonaktifkan input echo
    oldState, err := term.MakeRaw(fd)
    if err != nil {
        fmt.Println("failed to set raw mode:", err)
        return ""
    }

    defer term.Restore(fd, oldState)

    reader := bufio.NewReader(os.Stdin)
    var password []rune

    for {
        char, _, err := reader.ReadRune()
        if err != nil {
            break
        }

        // ENTER
        if char == '\n' || char == '\r' {
            fmt.Println()
            break
        }

        // BACKSPACE
        if char == 127 || char == 8 {
            if len(password) > 0 {
                password = password[:len(password)-1]
                fmt.Print("\b \b") // hapus 1 karakter di console
            }
            continue
        }

        // huruf / angka normal
        password = append(password, char)
        fmt.Print("*")
    }

    return string(password)
}


// ReadInt: parsing aman dari input → int
func ReadInt() int {
	for {
		input := ReadInput()
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid number, try again:")
			continue
		}
		return value
	}
}

// ReadFloat: parsing aman dari input → float64
func ReadFloat() float64 {
	for {
		input := ReadInput()
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid number, try again:")
			continue
		}
		return value
	}
}
