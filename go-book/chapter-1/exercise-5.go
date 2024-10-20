// read lines from the list of files and prints the count of lines that appears more than once.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	// fmt.Printf("no of args: %v", len(files))
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg) // returns 2 values - > *os.File, and built-in error type
			if err != nil {
				fmt.Fprintf(os.Stderr, "error while opening file: %v\n", err)
				continue
			}
			// fmt.Printf("opening file and starts reading it")
			countLines(f, counts)
			f.Close()
			fmt.Println(counts)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
