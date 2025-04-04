package main

import "fmt"

// main is the entry point of the application.
// It initializes a card using the newCard function and prints it.
func main() {
	// Use of explicit type declaration for clarity and readability.
	// Avoid using short variable declaration (:=) for package-level variables or when clarity is important.
	var card string = newCard()

	fmt.Println(card)
}

// newCard returns a string representing a playing card.
// This function can later be extended to return random or user-defined cards.
func newCard() string {
	return "Five of Diamonds"
}
