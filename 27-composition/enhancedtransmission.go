package main

import "fmt"

type EnhancedTransmission struct {
}

func (et EnhancedTransmission) ShiftUp() {
	fmt.Println("quick shift up......")
}

func (et EnhancedTransmission) ShiftDown() {
	fmt.Println("quick shift down......")
}
