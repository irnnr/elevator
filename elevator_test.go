package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPickupReturnsErrorForFloorPastMaxFloor(t *testing.T) {
	topFloor = 100
	el := Elevator{ Id:int32(1) }
	err := el.Pickup(1000, Down)

	assert.Error(t, err)
}

func TestPickupReturnsNilForFloorUpToMaxFloor(t *testing.T) {
	topFloor = 100
	el := Elevator{Id:int32(1)}
	err := el.Pickup(100, Down)

	assert.NoError(t, err)
}

func TestPickupReturnsNilForFloorAboveOrEqualFirstFloor(t *testing.T) {
	el := Elevator{Id:int32(1)}
	err := el.Pickup(1, Up)

	assert.NoError(t, err)
}

func TestPickupReturnsErrorForFloorBelowFirstFloor(t *testing.T) {
	el := Elevator{Id:int32(1)}
	err := el.Pickup(0, Up)

	assert.Error(t, err)
}

func TestStepMovesElevatorOneFloorDownForDirectionDown(t *testing.T) {
	el := Elevator{1, 100, 90, Down}
	el.Step()

	assert.EqualValues(t, 99, el.CurrentFloor)
}

func TestStepMovesElevatorOneFloorUpForDirectionUp(t *testing.T) {
	el := Elevator{1, 1, 90, Up}
	el.Step()

	assert.EqualValues(t, 2, el.CurrentFloor)
}