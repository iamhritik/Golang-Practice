package main

import "fmt"

func main() {
	colors := map[string]string{ // here key and values - string types
		"red":   "#ff0000", // all keys and value of same types
		"black": "#111111",
	}
	fmt.Println(colors)
	fmt.Println(colors["red"])
}
