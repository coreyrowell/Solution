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

    pointA := Point{X:0, Y:0}
    pointB := Point{X:50, Y:50}

    got := calculateDistance(pointA, pointB)
    want := 70.711

    if got != want {
        t.Errorf("got %f, wanted %f", got, want)
    }
}

func TestCalculateDistanceBetween3Points(t *testing.T) {
    pointA := Point{X:0, Y: 0}
    pointB := Point{X:50, Y:50}
    pointC := Point{X:100, Y:100}

    distanceToA := calculateDistance(pointA, pointB)
    distanceToB := calculateDistance(pointB, pointC)

    got := distanceToA + distanceToB
    want := 141.422

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

func TestAssignLoadsToDriver(t *testing.T) {
    runTest := func(t *testing.T, strategy LoadAssignmentStrategy, want [][]int) {
        load1 := Load{
            Number: 1,
            Start:  Point{X: -50.1, Y: 80.0},
            End:    Point{X: 90.1, Y: 12.2},
        }

        load2 := Load{
            Number: 2,
            Start:  Point{X: -24.5, Y: -19.2},
            End:    Point{X: 98.5, Y: 1.8},
        }

        load3 := Load{
            Number: 3,
            Start:  Point{X: 0.3, Y: 8.9},
            End:    Point{X: 40.9, Y: 55.0},
        }

        load4 := Load{
            Number: 4,
            Start:  Point{X: 5.3, Y: -61.1},
            End:    Point{X: 77.8, Y: -5.4},
        }

        loads := []Load{load1, load2, load3, load4}

        loadAssigner := &LoadAssigner{}
        loadAssigner.SetStrategy(strategy)

        got := loadAssigner.AssignLoads(loads)

        for i, driver := range got {
            for j, load := range driver {
                if len(driver) != len(want[i]) {
                    t.Errorf("mismatched route lengths")
                    return
                }
                if want[i][j] != load {
                    t.Errorf("mismatched load number")
                    return
                }
            }
        }
    }

    t.Run("Farthest Pickup Strategy", func(t *testing.T) {
        strategy := &FarthestPickupStrategy{}
        want := [][]int{
            {2, 1, 3},
            {4},
        }
        runTest(t, strategy, want)
    })

    t.Run("Nearest Pickup Strategy", func(t *testing.T) {
        strategy := &NearestPickupStrategy{}
        want := [][]int{
            {3, 2, 4},
            {1},
        }
        runTest(t, strategy, want)
    })

    t.Run("Unsorted Loads Pickup Strategy", func(t *testing.T) {
        strategy := &UnsortedLoadsPickupStrategy{}
        want := [][]int{
            {1, 2},
            {3, 4},
        }
        runTest(t, strategy, want)
    })
}