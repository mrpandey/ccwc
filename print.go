package main

import (
	"fmt"
	"sync"
)

// acquire this mutex before printing and using numWidth
var printMu sync.Mutex

// width of each number printed, must be a multiple of 8
var numWidth = 8

func GetNumsToPrint(count Count, opts Options) (numsToPrint []int) {
	if opts.PrintNewlineCount {
		numsToPrint = append(numsToPrint, count.Newlines)
	}

	if opts.PrintWordCount {
		numsToPrint = append(numsToPrint, count.Words)
	}

	if opts.PrintCharCount {
		numsToPrint = append(numsToPrint, count.Chars)
	}

	if opts.PrintByteCount {
		numsToPrint = append(numsToPrint, count.Bytes)
	}

	return numsToPrint
}

// Pass total count to this function to update numWidth
func UpdateNumWidth(c Count, opts Options) {
	nums := GetNumsToPrint(c, opts)

	max := 0
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	maxWidth := len(fmt.Sprintf("%d", max))

	printMu.Lock()

	if maxWidth > numWidth {
		// numWidth is next multiple of 8
		numWidth = maxWidth
		numWidth += 8 - maxWidth%8
	}

	printMu.Unlock()
}

func PrintOutput(nums []int, textName string) {
	// acquire lock for printing
	printMu.Lock()

	for _, num := range nums {
		fmt.Printf("%*d", numWidth, num)
	}

	fmt.Printf("\t%v\n", textName)

	printMu.Unlock()
}
