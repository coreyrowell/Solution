package strategies

import (
	"testing"
)

func TestUnsortedPickupStrategy(t *testing.T) {
    t.Run("Farthest Pickup Strategy", func(t *testing.T) {
        strategy := &UnsortedLoadsPickupStrategy{}
        want := [][]int{
            {1, 2},
            {3, 4},
        }
        runTest(t, strategy, want)
    })
}