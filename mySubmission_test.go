package main

import (
	"testing"
)

func TestPrecisionRoundedToTens(t *testing.T) {

    got := roundToDecimalPlaces(1.2345, 2)
    want := 1.23

    if got != want {
        t.Errorf("got %f, wanted %f", got, want)
    }
}

func TestPrecisionRoundedToHundreds(t *testing.T) {

    got := roundToDecimalPlaces(1.2345, 3)
    want := 1.235

    if got != want {
        t.Errorf("got %f, wanted %f", got, want)
    }
}

func TestCalculateDistanceBetween2Points(t *testing.T) {

    got := calculateDistance(0.0, 0.0, 50.0, 50.0)
    want := 70.71

    if got != want {
        t.Errorf("got %f, wanted %f", got, want)
    }
}

func TestCalculateDistanceBetween3Points(t *testing.T) {
    // @TODO - represent points as struct
    distanceToA := calculateDistance(0.0, 0.0, 50.0, 50.0)
    distanceToB := calculateDistance(50.0, 50.0, 100.0, 100.0)

    got := distanceToA + distanceToB
    want := 141.42

    if got != want {
        t.Errorf("got %f, wanted %f", got, want)
    }
}