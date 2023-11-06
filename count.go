package main

import (
	"unicode"
)

// Return number of bytes, characters, lines and words in the text.
type Counter func(text string, opts Options) (Count, error)

type Count struct {
	Bytes    int
	Chars    int
	Words    int
	Newlines int
}

// Counts the number of bytes, chars, words and newlines in a text and returns it.
var AllCounter = func(text string, _ Options) (Count, error) {
	numBytes := len(text)
	if numBytes == 0 {
		return Count{0, 0, 0, 0}, nil
	}

	numChars := 0
	numWords := 0
	numNewlines := 0
	isPartOfWord := false

	for _, rn := range text {
		numChars++

		if string(rn) == "\n" {
			numNewlines++
		}

		if unicode.IsSpace(rn) {
			isPartOfWord = false
		} else if !isPartOfWord {
			isPartOfWord = true
			numWords++
		}
	}

	return Count{
		Bytes:    numBytes,
		Chars:    numChars,
		Words:    numWords,
		Newlines: numNewlines,
	}, nil
}

func (c *Count) Add(cc Count) {
	c.Bytes += cc.Bytes
	c.Chars += cc.Chars
	c.Words += cc.Words
	c.Newlines += cc.Newlines
}
