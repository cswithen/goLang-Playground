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

	for _, link := range links {
		// the 'go' keyword sets up a new go routine which allows for the go scheduler to manage what code is running at what time.
		go checkLink(link)
	}
}

func checkLink(link string) {
	// typically we can pull both the response (resp) and the error (err) off of the http.Get() but in this case we are only looking to see if the Get request provides an error
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}
