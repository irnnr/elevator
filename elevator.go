package main


type Elevator struct {
	Id int32
	CurrentFloor int32
	GoalFloor int32
	Direction Direction
}

func (e Elevator) Status() (int32, int32, int32) {
	return e.Id, e.CurrentFloor, e.GoalFloor
}

func (e *Elevator) Pickup(floor int32, dir Direction) {
	e.GoalFloor = floor
	e.Direction = dir
}

func (e *Elevator) step() {

}