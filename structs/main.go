package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// // alex := person{"Hritik", "kumar"} // declaring and initializing a variable here
	// alex := person{firstName: "John", lastName: "Wick"} // one more way to use this sruct to declare struct
	// alex.firstName = "Hritik" // updating value here
	// fmt.Println(alex)
	// fmt.Println(alex.firstName)
	// fmt.Println(alex.lastName)

	///// other way to define struct here
	// var employee person // here firstName and lastName assigned empty string value
	// employee.firstName = "Hritik"
	// employee.lastName = "kumar"
	// fmt.Printf("%+v", employee) //it prints field namees along with their values | use for debugging & logging purpose

	//embedding struct
	jim := person{
		firstName: "Hritik",
		lastName:  "kumar",
		contact: contactInfo{
			email:   "hritik@github.com",
			zipCode: 11009999,
		},
	}
	fmt.Println(jim.contact.email)
	fmt.Printf("%+v", jim)
}
