package main

import (
	"fmt"
	"os"
)

// these can be reassigned to a mock of the corresponding function during testing
var wholeFileReader = WholeFileReader
var stdinReader = StdinReader
var allCounter = AllCounter

func Execute(parser Parser) {
	args := os.Args[1:]
	opts, textInputs, err := parser(args)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	totalCount := Count{}

	for _, txtIn := range textInputs {
		if txtIn.Type == StdIn {
			text, err := stdinReader()
			if err != nil {
				fmt.Println(err)
				continue
			}

			count, err := allCounter(text, opts)
			if err != nil {
				fmt.Println(err)
				continue
			}

			totalCount.Add(count)
			PrintOutput(count, opts, txtIn)

		} else if txtIn.Type == File {
			// TODO: use different strategy for large files
			text, err := wholeFileReader(txtIn.Name)
			if err != nil {
				fmt.Println(err)
				continue
			}

			count, err := allCounter(text, opts)
			if err != nil {
				fmt.Println(err)
				continue
			}

			totalCount.Add(count)
			PrintOutput(count, opts, txtIn)
		}
	}

	if len(textInputs) > 1 {
		PrintOutput(totalCount, opts, TextInput{Name: "total"})
	}

}
