package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in golang")
	rand.NewSource(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber{ // dicenumber has to match wit the case id 
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can Move to 3")
	case 4:
		fmt.Println("You can Move to 4")
	case 5:
		fmt.Println("You can Move to 5")
	case 6:
		fmt.Println("You can Move to 6")
	default:
		fmt.Println("You can Move to dat spot")
	}
}