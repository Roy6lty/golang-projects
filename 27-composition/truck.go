package main

import "fmt"

// embedding in Golang
type Truck struct {
	Engine
	Transmission
	SteeringWheel
}

func (t Truck) FourWheelDrive() {
	fmt.Println("Toggling 4WD......")
}
