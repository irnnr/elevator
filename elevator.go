package main

import (
	"errors"
	"fmt"
)


type Elevator struct {
	Id int32
	CurrentFloor int32
	GoalFloor int32
	Direction Direction
}

// Creates a new elevator going up from first floor
func NewElevator(id int32) Elevator {
	return Elevator{id, 1, 1, Up}
}

func (e Elevator) Status() (int32, int32, int32) {
	return e.Id, e.CurrentFloor, e.GoalFloor
}

func (e Elevator) IsBusy() bool {
	return (e.CurrentFloor != e.GoalFloor)
}

func (e *Elevator) Pickup(floor int32) error {
	if floor < 1 || floor > topFloor {
		return errors.New("Floor is above top floor")
	}

	e.GoalFloor = floor

	fmt.Printf("%d new goal: %d\n", e.Id, e.GoalFloor)

	return nil
}

func (e *Elevator) Step() {

	if e.CurrentFloor == e.GoalFloor {
		fmt.Printf("%d is at goal floor %d, not moving\n", e.Id, e.GoalFloor)
	} else if e.CurrentFloor < e.GoalFloor {
		e.up()
	} else {
		e.down()
	}
}

func (e *Elevator) up() {
	prevFloor := e.CurrentFloor

	if e.CurrentFloor != e.GoalFloor {
		e.CurrentFloor++
		fmt.Printf("%d going up: %d -> %d\n", e.Id, prevFloor, e.CurrentFloor)
	} else {
		fmt.Printf("%d is at goal floor %d, not moving\n", e.Id, e.GoalFloor)
	}
}

func (e *Elevator) down() {
	prevFloor := e.CurrentFloor

	if e.CurrentFloor != e.GoalFloor {
		e.CurrentFloor--
		fmt.Printf("%d going down: %d -> %d\n", e.Id, prevFloor, e.CurrentFloor)
	} else {
		fmt.Printf("%d is at goal floor %d, not moving\n", e.Id, e.GoalFloor)
	}
}