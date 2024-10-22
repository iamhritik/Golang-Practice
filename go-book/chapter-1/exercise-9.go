package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		var finalUrl string
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			finalUrl = "http://" + url
			fmt.Println(finalUrl)
		} else {
			finalUrl = url
		}
		resp, err := http.Get(finalUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
	}
}
