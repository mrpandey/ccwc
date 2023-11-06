package main

import "fmt"

func PrintOutput(count Count, opts Options, textIn TextInput) {
	if opts.PrintLineCount {
		fmt.Printf("%6d", count.Newlines)
	}

	if opts.PrintWordCount {
		fmt.Printf("%6d", count.Words)
	}

	if opts.PrintCharCount {
		fmt.Printf("%6d", count.Chars)
	}

	if opts.PrintByteCount {
		fmt.Printf("%5d", count.Bytes)
	}

	fmt.Printf("\t%v\n", textIn.Name)
}
