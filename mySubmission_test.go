package main

import (
	"testing"
)

// @TODO - it would be better to parametrize these test inputs

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
    // @TODO - use new Point struct
    distanceToA := calculateDistance(0.0, 0.0, 50.0, 50.0)
    distanceToB := calculateDistance(50.0, 50.0, 100.0, 100.0)

    got := distanceToA + distanceToB
    want := 141.42

    if got != want {
        t.Errorf("got %f, wanted %f", got, want)
    }
}

func TestParseLoadFromString(t *testing.T) {

    str := "200 (26.177425683364,117.72887590630494) (-49.79976850340818,59.79329676250279)"
    load := parseLoadFromString(str)
    wantLoadNumber := 200
    wantLoadStartX := 26.177425683364
    wantLoadStartY := 117.72887590630494
    wantLoadEndX := -49.79976850340818
    wantLoadEndY := 59.79329676250279

    if load.Number != wantLoadNumber {
        t.Errorf("got %d, wanted %d", load.Number, wantLoadNumber)
    }
    if load.Start.X != wantLoadStartX {
        t.Errorf("got %f, wanted %f", load.Start.X, wantLoadStartX)
    }
    if load.Start.Y != wantLoadStartY {
        t.Errorf("got %f, wanted %f", load.Start.X, wantLoadStartX)
    }
    if load.End.X != wantLoadEndX {
        t.Errorf("got %f, wanted %f", load.Start.X, wantLoadStartX)
    }
    if load.End.Y != wantLoadEndY {
        t.Errorf("got %f, wanted %f", load.Start.X, wantLoadStartX)
    }
}