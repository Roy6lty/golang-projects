package main

import "fmt"

type SteeringWheel struct {
}

func (sw SteeringWheel) TurnRight() {
	fmt.Println("Turning Left ......")
}

func (sw SteeringWheel) TurnLeft() {
	fmt.Println("Turning Right)...")

}
