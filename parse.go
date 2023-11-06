package main

import (
	"strings"
)

type Parser func(args []string) (Options, []TextInput, error)

type InputType string

const (
	File  InputType = "file"
	StdIn InputType = "stdin"
)

type Options struct {
	PrintByteCount bool
	PrintCharCount bool
	PrintWordCount bool
	PrintLineCount bool
}

type TextInput struct {
	Type InputType
	Name string
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
					opts.PrintLineCount = true
				default:
					return opts, inputs, NewInvalidOptionErr(char)
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

	if !(opts.PrintByteCount || opts.PrintCharCount || opts.PrintWordCount || opts.PrintLineCount) {
		opts.SetDefaultCountConfig()
	}

	return opts, inputs, nil
}

func (opts *Options) SetDefaultCountConfig() {
	opts.PrintByteCount = true
	opts.PrintLineCount = true
	opts.PrintWordCount = true
}
