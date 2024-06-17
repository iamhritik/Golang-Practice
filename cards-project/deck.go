package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck { // over here we are defining a func newDeck and it returns deck data type
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { // creating nested for loops over and '_' when we don't want to use that variable in the code
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() { // works as a receiver and instead of print you can define any method like set, get
	for i, card := range d {
		fmt.Println(i, card)
	}
}
