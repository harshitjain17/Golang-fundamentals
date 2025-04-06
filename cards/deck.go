package main

import "fmt"

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