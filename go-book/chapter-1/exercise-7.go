package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url) //fetch url using http.GET | resp.Body is of type io.ReadCloser - Reader => Read(p []byte) (n int, err error), Closter => Close() error
		if err != nil {
			fmt.Fprintf(os.Stderr, "URL Fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body) //func ReadAll(r Reader) ([]byte, error)
		resp.Body.Close()               //close
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
