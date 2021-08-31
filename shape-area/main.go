package main

import "fmt"

// interface declaration
type shape interface {
	getArea() float64
}

// triangle struct declaration
type triangle struct {
	base   float64
	height float64
}

// square struct declaration
type square struct {
	sideLength float64
}

func main() {
	// declaring and initializing the different structs
	tr := triangle{base: 10, height: 10}
	sq := square{sideLength: 10}

	// calling the interface function on the different structs that are associated to interface
	printArea(tr)
	printArea(sq)
}

// get area function on the triangle
func (t triangle) getArea() float64 {
	// fmt.Println(0.5 * base * height)
	return 0.5 * t.base * t.height
}

// get area function on the square
func (s square) getArea() float64 {
	// fmt.Println(sideLength * sideLength)
	return s.sideLength * s.sideLength
}

// interface function to print the return of the respective getArea function
func printArea(s shape) {
	fmt.Println(s.getArea())
}
