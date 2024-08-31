package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) //create a channel with string type - it's like any other var used in GO | chan - abbrevations of channel

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ { //integer based for  loop
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkLink(<-c, c) // blocking call - waiting for value from this <-c channel
	// }

	for l := range c { // assign the value that we got from this channel -> c
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
