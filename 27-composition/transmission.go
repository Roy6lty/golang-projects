package main

import "fmt"

type Transmission struct {
}

func (t Transmission) ShiftUp() {
	fmt.Println("Shift UP....")
}

func (t Transmission) ShiftDown() {
	fmt.Println("Shift Down....")

}
