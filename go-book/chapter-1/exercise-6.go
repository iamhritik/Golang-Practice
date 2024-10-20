package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	fileMap := make(map[string][]string) // key is string and value is slice of strings
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			fileMap[line] = append(fileMap[line], filename) // appending multiple values into slice of strings
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, strings.Join(fileMap[line], ", "))
		}
	}
}
