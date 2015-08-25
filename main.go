package main

import "fmt"

var elevators []Elevator
var numElevators = 2
var topFloor int32

func initElevators() {
	elevators = make([]Elevator, 0, numElevators)
	for i := 0; i < numElevators; i++ {
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