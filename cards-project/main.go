package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string { //need to define the return type of this func by adding string after ()
	return "Five of Diamonds"
}
