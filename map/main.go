package main

import "fmt"

func main() {
	// colors := map[string]string{ // here key and values - string types
	// 	"red":   "#ff0000", // all keys and value of same types
	// 	"black": "#111111",
	// }
	// fmt.Println(colors)
	// fmt.Println(colors["red"])

	// var colors map[string]string // declare a new variable but not assigned any value to it then GO automatically initialized it with 0 value - empty map
	// fmt.Println(colors) // prints - empty map -> map[]

	// colors := make(map[int]string) // creating map using make
	// colors[1] = "Hritik"
	// colors[2] = "Kumar"
	// delete(colors, 2) // delete keys and values from a map

	colors := map[string]string{
		"white": "#ff0000",
		"black": "#111111",
		"pink":  "#778787",
	}
	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
