package main

import (
	"fmt"
	"net/http"
	"time"
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

	// for { // it's not properly styled and getting clear understanding of this code
	// 	go checkLink(<-c, c) // blocking call - waiting for value from this <-c channel
	// }

	for l := range c { // range gets the value that we got from this channel -> c
		go func(link string) {
			// fmt.Println("channel range caller ------>")
			// time.Sleep(time.Second)     //time.Second -> is duration type -> int64
			time.Sleep(5 * time.Second) //if u want to sleep for 5sec -> simply multiply this time.Second with 5 => but this is a blocking call for MAIN Routine
			checkLink(link, c)
			// fmt.Println("called checkLink func() ------>")
		}(l) //invoking func literal (anonymous func) here | now this func literal not using l -> var value that's outer scope and can we changed by main routine
	}
}

func checkLink(link string, c chan string) {
	// time.Sleep(5 * time.Second) //if u want to sleep for 5sec -> simply multiply this time.Second with 5 => but this is a blocking call for MAIN Routine
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
