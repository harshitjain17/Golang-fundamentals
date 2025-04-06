package main

import "fmt"

func main() {
	// Create a new deck of cards
	cards := newDeck()

	// Print the deck of cards
	cards.print()

	// Deal a hand of 5 cards
	hand, remainingDeck := deal(cards, 5)
	fmt.Println("Hand:")
	hand.print()
	fmt.Println("Remaining Deck:")
	remainingDeck.print()

	// Save the deck to a file
	err := cards.saveToFile("my_cards.txt")
	if err != nil {
		fmt.Println("Error saving to file:", err)
	} else {
		fmt.Println("Deck saved to my_cards.txt")
	}

}