package main

import (
	"bufio"
	"io"
	"os"
	"unicode"
)

type Count struct {
	Bytes    int
	Chars    int
	Words    int
	Newlines int
}

func (c *Count) Add(other Count) {
	c.Bytes += other.Bytes
	c.Chars += other.Chars
	c.Words += other.Words
	c.Newlines += other.Newlines
}

func FileCounter(filename string, opts Options) (Count, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Count{}, ErrCannotOpenFile
	}

	defer file.Close()

	if opts.PrintByteCount && !opts.PrintCharCount && !opts.PrintWordCount && !opts.PrintLineCount {
		finfo, err := file.Stat()
		if err != nil {
			return Count{}, ErrCannotGetFileInfo
		}
		return Count{Bytes: int(finfo.Size())}, nil
	}

	return FullCounter(file)
}

func StdinCounter(opts Options) (Count, error) {
	if opts.PrintByteCount && !opts.PrintCharCount && !opts.PrintWordCount && !opts.PrintLineCount {
		return ByteCounter(os.Stdin)
	}

	return FullCounter(os.Stdin)
}

func FullCounter(r io.Reader) (Count, error) {
	reader := bufio.NewReader(r)

	count := Count{}
	isPartOfWord := false

	for {
		r, size, err := reader.ReadRune()
		// TODO: test case where a rune is split into multiple reads
		if err == io.EOF {
			break
		}

		if err != nil {
			return count, ErrCannotReadContent
		}

		if r == unicode.ReplacementChar && size == 1 {
			return count, ErrInvalidUTF8
		}

		count.Bytes += size
		count.Chars++

		if string(r) == "\n" {
			count.Newlines++
		}

		if unicode.IsSpace(r) {
			isPartOfWord = false
		} else if !isPartOfWord {
			isPartOfWord = true
			count.Words++
		}
	}

	return count, nil
}

func ByteCounter(r io.Reader) (Count, error) {
	reader := bufio.NewReader(r)
	// read 16kb at a time
	buf := make([]byte, 16*1024)
	count := Count{}

	for {
		n, err := reader.Read(buf)

		if err == io.EOF {
			count.Bytes += n
			break
		}

		if err != nil {
			return count, ErrCannotReadContent
		}

		count.Bytes += n
	}

	return count, nil
}
