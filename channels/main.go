// concurency is not parallelism
// channels are the only way different go routines can communicate to eachother
package main

import (
	"fmt"
	"net/http"
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

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

// input channel as a parameter
func checkLink(link string, c chan string) {
	// typically we can pull both the response (resp) and the error (err) off of the http.Get() but in this case we are only looking to see if the Get request provides an error
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, "is up!")
	c <- "Yep its up"
}
