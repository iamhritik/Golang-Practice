package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{} // custom interface

func main() {
	resp, err := http.Get("https://homelane.com")
	if err != nil {
		fmt.Println("Error is:", err)
		os.Exit(1)
	}

	//fmt.Println(resp) //right now it;s a HTTP header response but not the whole body
	// bs := make([]byte, 9999999) //creating byte slice with empty 9999 elements
	// resp.Body.Read(bs)          // Read takes bytes slice and read the data and put it in the byte slice
	// fmt.Println(string(bs))
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
