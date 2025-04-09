package main

import "fmt"

// contactInfo holds details like email and zip code
type contactInfo struct {
	email   string
	zipCode int
}

// person contains basic personal info along with contact details
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// Initializing a person struct using named fields
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@example.com",
			zipCode: 16802,
		},
	}

	jim.print()             // prints initial details
	jim.setName("Jimmy")    // updates firstName using a pointer receiver
	jim.print()             // prints updated details

	// Demonstrating behavior of slices (reference type)
	mySlice := []string{"Hi", "There", "how", "are", "you?"}
	updateSlice(mySlice)
	fmt.Println(mySlice)    // reflects changes made inside updateSlice
}

// setName updates the first name; uses pointer receiver to modify the actual struct
func (p *person) setName(newFirstName string) {
	p.firstName = newFirstName
}

// print outputs the contents of the person struct
func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// updateSlice modifies the contents of the passed-in slice (reference type)
func updateSlice(s []string) {
	s[0] = "Bye"
	s[1] = "There"
}
