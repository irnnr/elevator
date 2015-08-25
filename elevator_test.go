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