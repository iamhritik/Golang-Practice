package main

import (
	"fmt"
	"os"
)

func main() {
	for k, arg := range os.Args[1:] {
		fmt.Println(k, arg)
	}

}
