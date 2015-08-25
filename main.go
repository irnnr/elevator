package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func setGoals() {
	for _, el := range elevators {
		if !el.IsBusy() {
			// rand.Intn is zero based, ensure new goal is between 1 and topFloor
			newGoal := rand.Int31n(topFloor - 1) + 1

			newDir := Up
			randDown := rand.Intn(1)
			if randDown == 1 {
				newDir = Down
			}

			el.Pickup(newGoal, newDir)
		}

	}
}

func main() {
	topFloor = 100
	rand.Seed(time.Now().UTC().UnixNano())

	initElevators()
	globalStatus()

	setGoals()

}