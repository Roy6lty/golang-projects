package main

import "fmt"

func main(){
	fmt.Println("if else in golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular user"
	}else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number of even")
	} else {
		fmt.Println("Number of odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less then 10")
	} else {
		fmt.Println("Num is Not less than 10")
	}


	}
