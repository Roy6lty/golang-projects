package main

import "fmt"

func main() {
	fmt.Println("Welcome to fucntion in golang")
	greeter()

	result := adder(3,5)
	fmt.Println("Result is", result)

	proRes := proAdder(2,3,4,5,6)
	fmt.Println("result is", proRes)
	}

// adding a slice to a fucntion as varables
func proAdder(values...int)int{
	total := 0
	for _, val := range values {
		 total += val
	}
	return total
}


func adder(x int, y int) int {
	return x + y 

}



func greeterTwo(){
		fmt.Println("Another funtion")
	}

func greeter(){
	fmt.Println("Namstey from golang")
}