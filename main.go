package main

import "fmt"

var elevators []Elevator
var topFloor int32

func initElevators() {
	elevators = make([]Elevator, 0, 16)
	for i := 0; i < 16; i++ {
		elevators = append(elevators, Elevator{ Id:int32(i) })
	}
}

func globalStatus() {
	for _, el := range elevators {
		fmt.Println(el.Status())
	}
}

func main() {
	topFloor = 100

	initElevators()
	globalStatus()
}