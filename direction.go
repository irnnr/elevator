package main


type Direction struct {
	Direction string
}

var Up = Direction{ "up" }
var Down = Direction{ "down" }

func (d Direction) String() string {
	return d.Direction
}