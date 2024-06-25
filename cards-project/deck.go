package main

import ( // no need to use any separator like "," to separate these 2 packages that we are importing in this file
	"fmt"
	"os"
	"strings"
)

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

func deal(d deck, handSize int) (deck, deck) { //over here returning 2 values
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename) // err contains nil value if there's no error
	if err != nil {                  // error handling here
		// option #1 - log error and return a call to newDeck()
		// option #2 - log error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",") // splitting string using "," as separator
	return deck(s)
}
