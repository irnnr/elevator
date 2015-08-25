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

func (e *Elevator) Pickup(floor int32, dir Direction) error {
	if floor < 1 || floor > topFloor {
		return errors.New("Floor is above top floor")
	}

	e.GoalFloor = floor
	e.Direction = dir

	fmt.Printf("%d new goal: %s from %d\n", e.Id, dir, floor)

	return nil
}

func (e *Elevator) Step() {
	prevFloor := e.CurrentFloor

	switch e.Direction {
	case Up:
		e.up()
	case Down:
		e.down()
	}

	fmt.Printf("%d going %s : %d -> %d\n", e.Id, e.Direction, prevFloor, e.CurrentFloor)
}

func (e *Elevator) up() {
	e.CurrentFloor++
}

func (e *Elevator) down() {
	e.CurrentFloor--
}