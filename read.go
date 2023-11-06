package main

import (
	"bufio"
	"os"
)

// Reads entire text content from file at once.
// Not optimal for large files.
func WholeFileReader(filename string) (text string, err error) {
	textBytes, err := os.ReadFile(filename)

	if err != nil {
		return "", NewReadFileErr(filename)
	}

	return string(textBytes), nil
}

// Reads from stdin line by line and returns the text string.
func StdinReader() (text string, err error) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		text += line + "\n"
	}

	err = scanner.Err()
	if err != nil {
		return "", ErrCannotReadStdin
	}

	return text, nil
}
