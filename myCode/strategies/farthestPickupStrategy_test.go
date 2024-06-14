package strategies

import (
	"testing"
)

func TestFarthestPickupStrategy(t *testing.T) {
    t.Run("Farthest Pickup Strategy", func(t *testing.T) {
        strategy := &FarthestPickupStrategy{}
        want := [][]int{
            {2, 1, 3},
            {4},
        }
        runTest(t, strategy, want)
    })
}