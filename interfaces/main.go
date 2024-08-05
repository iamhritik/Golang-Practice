package main

import "fmt"

type bot interface { // everyone in this program - we have a new type - bot and if you are a type in this program with a func called "getGreeting()" and you return a string then you are now an honoary member of type bot and now you are an honoary member of this type bot - you can now call this func printGreeting() --> in this case englishtBot and spanishBot are also members of BOT type and can use this func - printGreeting()
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string { // receiver func removing eb from this - (eb englishBot) because we are using it in the func
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola !"
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) { // can declare 2 func with same even if the args value are differenct
// 	fmt.Println(sb.getGreeting())
// }
