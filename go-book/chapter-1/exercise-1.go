package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("testing")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // prints true - if there's a line and false where ther is no more input
		counts[input.Text()]++ //the lines read by input.Scan can be retrieved using input.Text()
	}
	// fmt.Println(counts)
	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, line)
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
