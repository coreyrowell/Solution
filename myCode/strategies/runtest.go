package strategies

import (
	"solution/myCode/load"
	"testing"
)

func runTest(t *testing.T, strategy load.LoadAssignmentStrategy, want [][]int) {
	load1 := load.Load{
		Number: 1,
		Start:  load.Point{X: -50.1, Y: 80.0},
		End:    load.Point{X: 90.1, Y: 12.2},
	}

	load2 := load.Load{
		Number: 2,
		Start:  load.Point{X: -24.5, Y: -19.2},
		End:    load.Point{X: 98.5, Y: 1.8},
	}

	load3 := load.Load{
		Number: 3,
		Start:  load.Point{X: 0.3, Y: 8.9},
		End:    load.Point{X: 40.9, Y: 55.0},
	}

	load4 := load.Load{
		Number: 4,
		Start:  load.Point{X: 5.3, Y: -61.1},
		End:    load.Point{X: 77.8, Y: -5.4},
	}

	loads := []load.Load{load1, load2, load3, load4}

	loadAssigner := &load.LoadAssigner{}
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