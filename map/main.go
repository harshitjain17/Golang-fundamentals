package main

import "fmt"

func main() {
	// Create a map to store the names and ages of people
	people := make(map[string]int)

	// Add some people to the map
	people["Alice"] = 30
	people["Bob"] = 25
	people["Charlie"] = 35

	// Print the map
	fmt.Println("People:", people)

	// Access a value in the map
	fmt.Println("Alice's age:", people["Alice"])

	// Check if a key exists in the map
	if age, exists := people["Bob"]; exists {
		fmt.Println(exists)
		fmt.Println("Bob's age:", age)
	} else {
		fmt.Println("Bob not found")
	}

	// Delete a key-value pair from the map
	delete(people, "Charlie")
	fmt.Println("After deleting Charlie:", people)

	// Iterate over the map and print each key-value pair
	for name, age := range people {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}