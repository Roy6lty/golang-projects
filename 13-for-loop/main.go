package main

import "fmt"

func main(){
	fmt.Println("Welcome to loops in go lang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	//set, condition, loop
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	
	    }

	for i := range days{
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("index is %v and is %v\n", index, day)
	}

	//using for as a while loop in python

	rougueValue := 1 
	for rougueValue < 10 {
		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}

	rougueValue2 := 1

	for rougueValue2 < 10 {

		if rougueValue2 == 2 {
			goto lco // jump out of for loop
		}

		if rougueValue == 5{
			rougueValue2++
			continue
		}

		fmt.Println("Value is:", rougueValue2)
		rougueValue2++

	}
	lco: //label
	fmt.Println("jumping at LearnCodeonLine.in")
}