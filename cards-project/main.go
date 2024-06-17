package main

func main() {
	//cards := []string{"Ace of Diamonds", newCard()} // type should be singular
	cards := deck{"Ace of Diamonds", newCard()} // here we are using deck type which is a slice of strings
	cards = append(cards, "Six of Spades")      // it will assign this value to this slice var and it's modifying the existing cards var but creating a new one

	// for i, card := range cards { // iterating cards slice var
	// 	fmt.Println(i, card) // print index and it's value
	// }
	cards.print()
}

func newCard() string { //need to define the return type of this func by adding string after ()
	return "Five of Diamonds"
}
