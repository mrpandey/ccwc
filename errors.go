package main

import "fmt"

var ErrCannotReadStdin = fmt.Errorf("invalid input")
var ErrCannotOpenFile = fmt.Errorf("cannot open file")
var ErrCannotReadFile = fmt.Errorf("cannot read file")
var ErrCannotReadContent = fmt.Errorf("failed to read content")
var ErrCannotGetFileInfo = fmt.Errorf("cannot get file info")
var ErrInvalidOption = fmt.Errorf("invalid option")
var ErrInvalidUTF8 = fmt.Errorf("invalid utf8")
