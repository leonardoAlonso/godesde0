package mypointers

import "fmt"

func UsePointer() {
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of pointer is, ", ptr)
	fmt.Println("Value of pointer is, ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value id, ", myNumber)
	fmt.Println("New value id, ", ptr)
	fmt.Println("New value id, ", *ptr)
}
