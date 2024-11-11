package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("welcome to video on slices")

	var fruitList = []string{"appple","Tomato","peach"} //intalization when created
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Bananna")
	fmt.Println(fruitList)

	fruitList = fruitList[:3]
	fmt.Println(fruitList)

	highscore := make([]int, 4) //craetes a pointer to a slice obj

	highscore[0] = 234
	highscore[1] = 345
	highscore[2] = 945
	highscore[3] = 867

	highscore = append(highscore, 155, 206, 377)

	sort.Ints(highscore)

	fmt.Println(highscore)

	// removing items from slices
	var courses = []string{"reactjs", "javascript", "swift"}
	courses = append(courses[:1], courses[2:]...)
	fmt.Println(courses)

}
