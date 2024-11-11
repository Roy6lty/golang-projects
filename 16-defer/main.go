package main

import "fmt"

func main(){
	defer fmt.Println("world")
	defer fmt.Println("hello")
	fmt.Println("one")
	fmt.Println("two")

	myDefer()
}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}