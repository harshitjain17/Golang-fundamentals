package main

import (
	"os"
	"testing"
)

// Test the newDeck function to ensure it creates a deck of 52 cards
func TestNewDeck(t *testing.T) {
	
	// Create a new deck of cards
	cards := newDeck()

	// Check the length of the deck
	if len(cards) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(cards))
	}

	// Check the first card in the deck
	expectedFirstCard := "Ace of Spades"
	if cards[0] != expectedFirstCard {
		t.Errorf("Expected first card to be '%s', but got '%s'", expectedFirstCard, cards[0])
	}

	// Check the last card in the deck
	expectedLastCard := "King of Clubs"
	if cards[len(cards)-1] != expectedLastCard {
		t.Errorf("Expected last card to be '%s', but got '%s'", expectedLastCard, cards[len(cards)-1])
	}
}

// Test the saveToFile and newDeckFromFile functions to ensure they work correctly
func TestSavetoFileAndNewDeckFromFile(t *testing.T) {

	// Remove the test file if it exists
	_ = os.Remove("_decktesting")
	
	// Create a new deck of cards
	cards := newDeck()

	// Save the deck to a file
	err := cards.saveToFile("_decktesting")
	if err != nil {
		t.Errorf("Error saving to file: %v", err)
	}

	// Load the deck from the file
	loadedDeck, err := newDeckFromFile("_decktesting")
	if err != nil {
		t.Errorf("Error loading from file: %v", err)
	}

	// Check if the loaded deck is equal to the original deck
	if len(loadedDeck) != len(cards) {
		t.Errorf("Expected loaded deck length of %v, but got %v", len(cards), len(loadedDeck))
	}

	// Check if the cards in the loaded deck match the original deck
	for i, card := range cards {
		if card != loadedDeck[i] {
			t.Errorf("Expected card %d to be '%s', but got '%s'", i, card, loadedDeck[i])
		}
	}

	// Remove the test file after the test
	_ = os.Remove("_decktesting")
}