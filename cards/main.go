package main

import "fmt"

func main() {
	// Create a slice of cards and initialize with a few values
	var cards = []string{newCard(), "Two of Spades", "Ace of Diamonds"}

	// Append a new card to the slice
	cards = append(cards, "Three of Hearts")

	// Loop through the cards and print each one with its index
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// newCard returns a predefined card as a string
func newCard() string {
	return "Five of Diamonds"
}
