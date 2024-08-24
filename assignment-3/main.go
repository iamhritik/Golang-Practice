package main

import (
	"fmt"
	"io"
	"os"
)

type dataWriter struct{}

func main() {
	fmt.Println(os.Args[1])

	file, err := os.Open(os.Args[1]) // For read access.
	if err != nil {
		fmt.Println("Error is: ", err)
		os.Exit(1)
	}
	// // over here we are creating a BYTE slice and then calling READ() function to read the actual data and put it into this BYTE Slice and then we are printing that data using fmt.Println()
	// bs := make([]byte, 999999)
	// file.Read(bs)
	// fmt.Println(string(bs))

	// using io.Copy with os.Stdout to print the file data on the terminal
	// io.Copy(os.Stdout, file)

	// over here we are using custom Write() func to print the data on the terminal
	dr := dataWriter{}
	io.Copy(dr, file)
}

func (dataWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Successfully write all the data on the Terminal ----")
	return len(bs), nil
}
