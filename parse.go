package main

import (
	"fmt"
	"strings"
)

type Parser func(args []string) (Options, []TextInput, error)

type InputType string

const (
	File  InputType = "file"
	StdIn InputType = "stdin"
)

type Options struct {
	PrintByteCount    bool
	PrintCharCount    bool
	PrintWordCount    bool
	PrintNewlineCount bool
}

type TextInput struct {
	Type  InputType
	Name  string
	Index int
}

var DefaultParser = func(args []string) (Options, []TextInput, error) {
	opts := Options{}
	inputs := []TextInput{}

	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			// parse cli options

			if arg == "-" {
				inputs = append(inputs, TextInput{
					Type: StdIn,
					Name: arg,
				})

				continue
			}

			arg = arg[1:]

			// TODO: add help option
			for _, rn := range arg {
				char := string(rn)

				switch char {
				case "c":
					opts.PrintByteCount = true
				case "m":
					opts.PrintCharCount = true
				case "w":
					opts.PrintWordCount = true
				case "l":
					opts.PrintNewlineCount = true
				default:
					return opts, inputs, fmt.Errorf("%w: %v", ErrInvalidOption, char)
				}
			}
		} else {
			// arg is a filename
			inputs = append(inputs, TextInput{
				Type: File,
				Name: arg,
			})
		}
	}

	if len(inputs) == 0 {
		// read from stdin
		inputs = append(inputs, TextInput{
			Type: StdIn,
			Name: "",
		})
	}

	for i := range inputs {
		inputs[i].Index = i
	}

	if !(opts.PrintByteCount || opts.PrintCharCount || opts.PrintWordCount || opts.PrintNewlineCount) {
		opts.SetDefault()
	}

	return opts, inputs, nil
}

func (opts *Options) SetDefault() {
	opts.PrintByteCount = true
	opts.PrintNewlineCount = true
	opts.PrintWordCount = true
}
