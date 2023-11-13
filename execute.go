package main

import (
	"fmt"
	"os"
)

// these can be reassigned to a mock of the corresponding function during testing
var stdinCounter = StdinCounter
var fileCounter = FileCounter

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
			count, err := stdinCounter(opts)
			if err != nil {
				fmt.Printf("%v: %v\n", txtIn.Name, err)
				continue
			}

			totalCount.Add(count)
			PrintOutput(count, opts, txtIn)

		} else if txtIn.Type == File {
			count, err := fileCounter(txtIn.Name, opts)
			if err != nil {
				fmt.Printf("%v: %v\n", txtIn.Name, err)
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
