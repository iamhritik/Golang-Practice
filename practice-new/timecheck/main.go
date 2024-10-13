package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println("current time is: ", presentTime)
	fmt.Println("current time format is", presentTime.Format("01-02-2006")) //always need to use this specific format - 01-02-2006
	fmt.Println("current day is", presentTime.Format("Monday"))
	fmt.Println("current time is", presentTime.Format("15:04:05")) ///always need to use this specific format - 15:04:05
}
