package main

import (
	"strconv"
	"testing"
)

func TestCalculateDistance(t *testing.T) {

    got := true
    want := true

    if got != want {
        fgot := strconv.FormatBool(got)
        fwant := strconv.FormatBool(want)
        t.Errorf("got %q, wanted %q", fgot, fwant)
    }
}