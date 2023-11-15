package main

import (
	"fmt"
	"os"
	"sync"
)

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

	wg := sync.WaitGroup{}
	channels := make([]chan *InputCount, lenInput)

	stdInputs := []TextInput{}

	for _, txtIn := range textInputs {
		channels[txtIn.Index] = make(chan *InputCount)

		if txtIn.Type == StdIn {
			stdInputs = append(stdInputs, txtIn)
		} else {
			wg.Add(1)
			go ProcessFileInput(txtIn, opts, channels[txtIn.Index], &wg)
		}
	}

	// process all stdin sequentially in a separate go routine
	wg.Add(1)
	go ProcessStdInput(stdInputs, opts, channels, &wg)

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

func ProcessFileInput(input TextInput, opts Options, ch chan *InputCount, wg *sync.WaitGroup) {
	defer wg.Done()

	count, err := fileCounter(input.Name, opts)

	if err != nil {
		printMu.Lock()
		fmt.Printf("%v: %v", input.Name, err)
		printMu.Unlock()

		ch <- nil
	} else {
		ch <- &InputCount{Count: count, Input: input}
	}
}

func ProcessStdInput(inputs []TextInput, opts Options, chs []chan *InputCount, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, input := range inputs {
		count, err := stdinCounter(opts)

		if err != nil {
			printMu.Lock()
			fmt.Printf("stdin: %v", err)
			printMu.Unlock()

			chs[input.Index] <- nil
		} else {
			chs[input.Index] <- &InputCount{Count: count, Input: input}
		}
	}

}
