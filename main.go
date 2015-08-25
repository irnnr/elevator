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
		elevators = append(elevators, NewElevator(int32(i)))
	}
}

func globalStatus() {
	for _, el := range elevators {
		fmt.Println(el.Status())
	}
}

func setGoals() {
	for i := range elevators {
		if !elevators[i].IsBusy() {
			// rand.Intn is zero based, ensure new goal is between 1 and topFloor
			newGoal := rand.Int31n(topFloor - 1) + 1

			newDir := Up
			randDown := rand.Intn(1)
			if randDown == 1 {
				newDir = Down
			}

			elevators[i].Pickup(newGoal, newDir)
		}

	}
}

func moveElevators() {
	for i := range elevators {
		elevators[i].Step()
	}
}

func main() {
	topFloor = 100
	rand.Seed(time.Now().UTC().UnixNano())

	initElevators()
	globalStatus()

	setGoals()


	for {
		moveElevators()
		time.Sleep(time.Second)
	}

}