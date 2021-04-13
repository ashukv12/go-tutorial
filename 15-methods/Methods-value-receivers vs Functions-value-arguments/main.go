package main

import (
	"fmt"
)

// Struct type - `Point`
type Point struct {
	X, Y float64
}

func (p Point) IsAbove(y float64) bool {
	return p.Y > y
}

func IsAboveFunc(p Point, y float64) bool {
	return p.Y > y
}

func main() {
	p := Point{0, 4}
	ptr := &p

	fmt.Println("Point p = ", p)

	// Calling a Method with Value receiver
	fmt.Println(p.IsAbove(1))   // Valid
	fmt.Println(ptr.IsAbove(1)) // Valid

	// Calling a Function with a Value argument
	fmt.Println(IsAboveFunc(p, 1)) // Valid
	// Cannot use ptr (type *Point) as type Point in argument to IsAboveFunc
	// IsAboveFunc(ptr, 1) // Not Valid
}
