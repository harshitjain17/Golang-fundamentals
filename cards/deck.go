package main

import "fmt"

//  Create a new type of 'deck' which is a slice of strings
type deck []string

func (d deck) print() {
	// Loop through the cards and print each one with its index
	for i, card := range d {
		fmt.Println(i, card)
	}
}