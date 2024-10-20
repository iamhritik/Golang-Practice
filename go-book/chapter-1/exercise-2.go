package main

import (
	"fmt"
	"os"
	"strings"
)

// var s, sep string = "", ""
var s, sep string

// func main() {
// 	// s, sep := "", ""
// 	for _, arg := range os.Args[1:] {
// 		sep += s + arg
// 		s = " "
// 	}
// 	fmt.Println(sep)
// }

// option -2 using - strings.join

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
