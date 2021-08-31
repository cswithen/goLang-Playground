package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
// slice is an adjustable array

type deck []string

// need to declare the exact datatype that is returned from the function
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// replace unused but required variables with and underscore_
	// for i, suit := range cardSuits
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// funciton deal has a paramater (d with type deck) and handsize with type int and returns two decks
func deal(d deck, handSize int) (deck, deck) {

	//return from the beginning of the deck to the handSize, and then the remainder of the deck
	return d[:handSize], d[handSize:]
}

// function toString which is a reciever function to (d with type deck) returns a string
func (d deck) toString() string {

	//from the strings default package, (sliceOfstrings, separator string) (https://pkg.go.dev/strings@go1.17#Join)
	return strings.Join([]string(d), ",")

}

// function saveToFile is a reciever function attached to (d with type deck) and uses a parameter called filename returns a value of type error in case it hits one
func (d deck) saveToFile(filename string) error {

	// ioutil is a default package and we are using the WriteFile function (https://pkg.go.dev/io/ioutil@go1.17#example-WriteFile), (filename, []byte, permissions)
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// function newDeckFromFile uses (filename type string) as a parameter and returns something with the type deck
func newDeckFromFile(filename string) deck {
	// we return two things from the ReadFile function (https://pkg.go.dev/io/ioutil@go1.17#example-ReadFile)
	// first is a byteSlice (bs) and second is an error (err)
	// nill means no value
	bs, err := ioutil.ReadFile(filename)

	// this is (! = => !=)
	if err != nil {
		// Log the error and entirely quit the program
		fmt.Println("Error:", err)

		// (https://pkg.go.dev/os@go1.17#Exit) this will exit the program entirely
		os.Exit(1)
	}

	// make the bytes a string and then split into a slice using Slice(https://pkg.go.dev/strings@go1.17#Split)
	s := strings.Split(string(bs), ",")
	return deck(s)
}

// function to shuffle the deck it is called on, it is a recievier funciton (d deck) with no return
func (d deck) shuffle() {

	//random number generator that will create a new seed/source (https://pkg.go.dev/math/rand@go1.17#Source) using the current time (https://pkg.go.dev/time@go1.17#Now)
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// i is index of the card, no need to call the one card
	for i := range d {
		// len = length; Intn (https://pkg.go.dev/math/rand@go1.17#Intn)
		// randon generator uses a seed generator (https://pkg.go.dev/math/rand@go1.17#Rand)
		newPosition := r.Intn(len(d) - 1)

		//take whatever is at i and swap with newPosition and vice versa
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
