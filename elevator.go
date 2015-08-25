package main

import (
	"errors"
)


type Elevator struct {
	Id int32
	CurrentFloor int32
	GoalFloor int32
	Direction Direction
}

func (e Elevator) Status() (int32, int32, int32) {
	return e.Id, e.CurrentFloor, e.GoalFloor
}

func (e *Elevator) Pickup(floor int32, dir Direction) error {
	if floor < 1 || floor > topFloor {
		return errors.New("Floor is above top floor")
	}

	e.GoalFloor = floor
	e.Direction = dir

	return nil
}

func (e *Elevator) step() {

}