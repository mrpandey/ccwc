package main

import "fmt"

var ErrCannotReadStdin = fmt.Errorf("invalid input")

func NewReadFileErr(filename string) error {
	return fmt.Errorf("cannot read file %v", filename)
}

func NewInvalidOptionErr(char string) error {
	return fmt.Errorf("invalid option '%v'", char)
}
