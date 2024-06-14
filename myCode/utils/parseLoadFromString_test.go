package utils

import (
	"testing"
)

func TestParseLoadFromString(t *testing.T) {

    str := "200 (26.177425683364,117.72887590630494) (-49.79976850340818,59.79329676250279)"
    load := ParseLoadFromString(str)
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