package main

import "fmt"

func main() {
	fmt.Println("welcome to class on pointers")

	// var ptr  *int  -- points to a location in memory
	// fmt.Println("value of pointer is", ptr)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("value of actual poniter is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New Value is:", myNumber)
} 