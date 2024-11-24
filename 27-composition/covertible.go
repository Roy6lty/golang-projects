package main

import "fmt"

type transmission interface {
	ShiftUp()
	ShiftDown()
}
type Convertable struct {
	Engine
	transmission
	SteeringWheel
}

func (c *Convertable) CoverTop() {
	fmt.Println("Covering top...")

}
