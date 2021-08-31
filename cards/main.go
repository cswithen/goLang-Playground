//* To run these files 'go run main.go deck.go'
package main

func main() {
	// can ommit 'string' as it could be inferred from the other side
	//   longform way of declaring vairables
	// var card string = "Ace of Spades"

	// ":=" only used for defining a new variable, just "=" afterwards
	cards := newDeck()

	//no need to import here because it is declared in the same passage
	//hand, remainingDeck := <deck>, <deck>
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
	// cards.print()

	// this will print out remainingDeck as one string
	remainingDeck.toString()

	//this will save the deck cards to a file in this folder
	cards.saveToFile("myCards")

	//this will open the saved cards
	pulledCards := newDeckFromFile("myCards")

	//this will shuffle a deck of cards
	pulledCards.shuffle()
	pulledCards.print()

}
