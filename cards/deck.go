package main

import (
	"fmt"
	"os"
	"strings"
)

//  Create a new type of 'deck' which is a slice of strings
type deck []string

// Create a new deck of cards with 52 cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Loop through the cards and print each one with its index
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal a hand of cards from the deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Convert the deck to a string representation
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save the deck to a file with the given filename
func (d deck) saveToFile(filename string) error {
	data := d.toString()
	return os.WriteFile(filename, []byte(data), 0666)
}

// Load a deck from a file with the given filename
func newDeckFromFile(filename string) (deck, error) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	} 
	return deck(strings.Split(string(bs), ",")), nil
}