package strategies

import (
	"testing"
)

func TestNearestPickupStrategy(t *testing.T) {
    t.Run("Farthest Pickup Strategy", func(t *testing.T) {
        strategy := &NearestPickupStrategy{}
        want := [][]int{
            {3, 2, 4},
            {1},
        }
        runTest(t, strategy, want)
    })
}