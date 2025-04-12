package main

import "fmt"

// shape defines a common interface for geometric shapes.
type shape interface {
	getArea() float64
}

// triangle represents a triangle with height and base.
type triangle struct {
	height float64
	base   float64
}

// square represents a square with a given side length.
type square struct {
	sideLength float64
}

func main() {
	t1 := triangle{
		height: 2,
		base:   3,
	}

	s1 := square{
		sideLength: 4,
	}

	printArea(t1)
	printArea(s1)
}

// getArea calculates the area of a triangle.
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

// getArea calculates the area of a square.
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

// printArea prints the area of any shape that implements the shape interface.
func printArea(s shape) {
	fmt.Println(s.getArea())
}
