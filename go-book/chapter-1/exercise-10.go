package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
		fmt.Printf("URL status code is: %d\n", resp.StatusCode)
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
	}
}
