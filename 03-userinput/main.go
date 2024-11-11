package main

import (
	"bufio"
	"fmt"
	"os"
)

func main (){
	welcome := "welcome to user input"
	fmt.Println(welcome )

	reader := bufio.NewReader(os.Stdin) // standard input 
	fmt.Println("Enter the rating for our Pizza")

	//comma ok || err ok

	input, _ := reader.ReadString('\n') // return a tuple of a string and  if any error 
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("format for rating %T", input)

}

