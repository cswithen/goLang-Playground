package main

import "fmt"

// interface that will allow us to reuse logic for acts kindof like js prototypes
type bot interface {
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

// GET GREETING FUNCS
// can remove the eb if we are not actually using it
func (eb englishBot) getGreeting() string {
	// custom logic for generating an english greeting
	return "Hi There!"
}

// can remove the sb if we are not actually using it
func (sb spanishBot) getGreeting() string {
	// custom logic for generating an spanish greeting
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
