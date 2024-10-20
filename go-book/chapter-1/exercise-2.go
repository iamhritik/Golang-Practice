package main

import (
	"fmt"
	"os"
)

// var s, sep string = "", ""
var s, sep string

func main() {
	// s, sep := "", ""
	for _, arg := range os.Args[1:] {
		sep += s + arg
		s = " "
	}
	fmt.Println(sep)
}
