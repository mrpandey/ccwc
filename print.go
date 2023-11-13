package main

import "fmt"

func PrintOutput(count Count, opts Options, textIn TextInput) {
	formatStr := "%*d"
	numWidth := getNumWidth([]int{count.Newlines, count.Words, count.Chars, count.Bytes})
	numWidth += 4

	if opts.PrintLineCount {
		fmt.Printf(formatStr, numWidth, count.Newlines)
	}

	if opts.PrintWordCount {
		fmt.Printf(formatStr, numWidth, count.Words)
	}

	if opts.PrintCharCount {
		fmt.Printf(formatStr, numWidth, count.Chars)
	}

	if opts.PrintByteCount {
		fmt.Printf(formatStr, numWidth, count.Bytes)
	}

	fmt.Printf("\t%v\n", textIn.Name)
}

func getNumWidth(nums []int) int {
	max := 0

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return len(fmt.Sprintf("%d", max))
}
