package main

import "fmt"

func main(){
	fmt.Println("welcome to arrays in golangs")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Pineapple"

	fmt.Println("Fruits List is:", fruitList) 
	fmt.Println("Fruits list is", len(fruitList)) // retruns 4

	vegList := [5]string{"potato","beans","mushroom"}
	fmt.Println("Vegy list is:", len(vegList)) //


}