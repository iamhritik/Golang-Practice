package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of Spades"   // over here we are relying on compiler to find out the data type of this variable by using := | we are using := Operator only when defining a new variable, not when reassigning value to a var | also worked only within the func block
	card = "Five of Diamonds" // reassigning values here
	fmt.Println(card)
}
