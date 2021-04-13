package main

import "fmt"

func main() {
	var p *int
	fmt.Println("p = ", p)

	// Initializing
	var a = 5.67
	var pt = &a

	fmt.Println("Value stored in variable a = ", a)
	fmt.Println("Address of variable a = ", &a)
	fmt.Println("Value stored in variable p = ", pt)

	// Creation
	ptr := new(int) // Pointer to an `int` type
	*ptr = 100

	fmt.Printf("Ptr = %#x, Ptr value = %d\n", ptr, *ptr)

}
