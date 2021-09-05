// concurency is not parallelism
// channels are the only way different go routines can communicate to eachother
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// setting a slice of strings for a bunch of websites
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// time to create a channel
	c := make(chan string)
	// channels are used to pass data between routines
	// syntax is :
	// channel <- 5 // to pass the number 5 into the channel
	// myNumber <- channel // to pull the number from the channel and assign it to a variable
	// fmt.Println(<-channel) // advanced syntax to print something once received in the channel w/o saving the variable

	for _, link := range links {
		// the 'go' keyword sets up a new go routine which allows for the go scheduler to manage what code is running at what time.
		go checkLink(link, c)
	}
	// this will tell the channel to wait for !a! call and then stop the main routine
	// fmt.Println(<-c)

	// this is a for loop that runs for the length of the links slice
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	/*
		// this is a for loop that will continue to run until the program is finished (infinite loop)
		for {
			// this repeats the creation of the go routine
			go checkLink(<-c, c)
		}
	*/

	/*
		// refactor of the loop above to make it 'easier' to debug
		// wait for the channel to receive a value and then assign it to the variable l
		for l := range c {
			// pauses the current go routine for x number of seconds time.Sleep(x * time.Second)
			time.Sleep(5 * time.Second)
			go checkLink(l, c)
		}
	*/

	// this defines a function literal[anonymous function] the () at the end calls the function literal
	for l := range c {
		// in order to utilize the pass by value nuance in go
		// pass in l into the 'function literal' as an argument
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
			// and then when calling the function be sure to use l
		}(l)
	}
}

// input channel as a parameter
func checkLink(link string, c chan string) {
	// typically we can pull both the response (resp) and the error (err) off of the http.Get() but in this case we are only looking to see if the Get request provides an error
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// this pushes a message into the channel
		// c <- "Might be down I think"
		// this pushes the link variable into the channel
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	// this pushes a message into the channel
	// c <- "Yep its up"
	// this pushes the link variable into the channel
	c <- link
}
