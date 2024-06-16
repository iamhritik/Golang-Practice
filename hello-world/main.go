package main // what does package main mean ? -> Package is like a project or workspace | collections of code files| requirement of package named main is that we define a function named MAIN that runs automatically when the program runs

import "fmt" //what does import fmt mean ? -> It gives acccess to all the code that exists in this fmt (part of standard library)

func main() { //what's this func main() ->
	fmt.Println("Hello World")
}

// how do we run the code in our project ? -> go run main.go
// how is the main.go file organized ? -> Package Declaratiom => Import other packages that we need => Declare functions, tell Go to do things
