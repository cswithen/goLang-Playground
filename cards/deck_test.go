package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	//test to see the deck length
	if len(d) != 52 {
		// Errorf is formatted error where the %v injects the argument (len(d))
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	//test to see what the first element of the deck is
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	//test to see what the last element of the deck is
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first card of King of Clubs but got %v", d[len(d)-1])
	}

}

//test to see if we can save and load a deck
func TestSaveDeckAndNewDeckFromFile(t *testing.T) {
	// delete any file with the name _decktesting (https://pkg.go.dev/os@go1.17#Remove)
	os.Remove("_decktesting")

	//create a new deck and save to file
	deck := newDeck()
	deck.saveToFile("_decktesting")

	//load deck from file
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
