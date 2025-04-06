package main

func main() {
	// Create a slice of cards and initialize with a few values
	var cards = deck{newCard(), "Two of Spades", "Ace of Diamonds"}

	// Append a new card to the slice
	cards = append(cards, "Three of Hearts")

	cards.print()
}

// newCard returns a predefined card as a string
func newCard() string {
	return "Five of Diamonds"
}
