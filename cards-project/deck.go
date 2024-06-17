package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings

type deck []string

func (d deck) print() { // works as a receiver and instead of print you can define any method like set, get
	for i, card := range d {
		fmt.Println(i, card)
	}
}
