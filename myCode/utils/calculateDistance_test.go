package utils

import (
	"solution/myCode/load"
	"testing"
)

func TestCalculateDistanceBetween2Points(t *testing.T) {

    pointA := load.Point{X:0, Y:0}
    pointB := load.Point{X:50, Y:50}

    got := CalculateDistance(pointA, pointB)
    want := 70.711

    if got != want {
        t.Errorf("got %f, wanted %f", got, want)
    }
}

func TestCalculateDistanceBetween3Points(t *testing.T) {
    pointA := load.Point{X:0, Y: 0}
    pointB := load.Point{X:50, Y:50}
    pointC := load.Point{X:100, Y:100}

    distanceToA := CalculateDistance(pointA, pointB)
    distanceToB := CalculateDistance(pointB, pointC)

    got := distanceToA + distanceToB
    want := 141.422

    if got != want {
        t.Errorf("got %f, wanted %f", got, want)
    }
}