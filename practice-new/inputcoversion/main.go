package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter your ratings out of 5")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //changing string into float64
	if err != nil {
		fmt.Println("some error:", err)
	} else {
		fmt.Println("ratings are: ", numRating)
		fmt.Printf("type of rating is %T", numRating)
	}
}
