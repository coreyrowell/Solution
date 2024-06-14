package utils

import (
	"testing"
)

func TestPrecisionRoundedToTens(t *testing.T) {

    got := RoundToDecimalPlaces(1.2345, 2)
    want := 1.23

    if got != want {
        t.Errorf("got %f, wanted %f", got, want)
    }
}

func TestPrecisionRoundedToHundreds(t *testing.T) {

    got := RoundToDecimalPlaces(1.2345, 3)
    want := 1.235

    if got != want {
        t.Errorf("got %f, wanted %f", got, want)
    }
}