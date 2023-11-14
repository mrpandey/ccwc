package main

import (
	"fmt"
	"os"
	"sync"
)

// acquire this mutex before printing and using numWidth
var printMu sync.Mutex

// width of each number printed, must be a multiple of 8
var numWidth = 8

// these can be reassigned to a mock of the corresponding function during testing
var stdinCounter = StdinCounter
var fileCounter = FileCounter

type InputCount struct {
	Input TextInput
	Count Count
}

func Execute(parser Parser) {
	args := os.Args[1:]
	opts, textInputs, err := parser(args)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lenInput := len(textInputs)

	totalCount := Count{}

	channels := make([]chan *InputCount, lenInput)
	wg := sync.WaitGroup{}

	for i, txtIn := range textInputs {
		channels[i] = make(chan *InputCount)
		wg.Add(1)
		go ProcessInput(txtIn, opts, channels[i], &wg)
	}

	// print result in the same order of textInputs in a separate go routine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range channels {
			inputCount := <-channels[i]
			if inputCount == nil {
				continue
			}

			totalCount.Add(inputCount.Count)
			UpdateNumWidth(totalCount, opts)

			nums := GetNumsToPrint(inputCount.Count, opts)
			PrintOutput(nums, inputCount.Input.Name)
		}
	}()

	wg.Wait()

	if lenInput > 1 {
		nums := GetNumsToPrint(totalCount, opts)
		PrintOutput(nums, "total")
	}
}

func ProcessInput(input TextInput, opts Options, ch chan *InputCount, wg *sync.WaitGroup) {
	defer wg.Done()

	var err error
	var count Count

	if input.Type == StdIn {
		count, err = stdinCounter(opts)
	} else {
		count, err = fileCounter(input.Name, opts)
	}

	if err != nil {
		printMu.Lock()
		fmt.Printf("%v: %v", input.Name, err)
		printMu.Unlock()

		ch <- nil
	} else {
		ch <- &InputCount{Count: count, Input: input}
	}
}
