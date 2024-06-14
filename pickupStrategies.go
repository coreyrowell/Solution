package main

import "sort"


type FarthestPickupStrategy struct{}

func (s *FarthestPickupStrategy) AssignLoadsToDriver(loads []Load) [][]int {
    drivers := [][]int{}
    remainingLoads := make([]Load, len(loads))
    copy(remainingLoads, loads)

    for len(remainingLoads) > 0 {
        driverLoads := []int{}
        currentTime := 0.0
        currentPosition := Point{X: 0, Y: 0}

        // Sort remainingLoads such that loads with endpoints further from current position come later in list
        sort.Slice(remainingLoads, func(i, j int) bool {
			distI := calculateDistance(currentPosition, remainingLoads[i].End)
			distJ := calculateDistance(currentPosition, remainingLoads[j].End)
			return distI > distJ
		})

        for _, load := range remainingLoads {
            loadNumber := load.Number
            pickup := load.Start
            dropoff := load.End
            timeToPickup := calculateDistance(currentPosition, pickup)
            timeToDropoff := calculateDistance(pickup, dropoff)
            timeToDepot := calculateDistance(dropoff, Point{X:0, Y:0})

            if (currentTime + timeToPickup + timeToDropoff + timeToDepot) <= 720 {
                driverLoads = append(driverLoads, loadNumber)
                currentTime += timeToPickup + timeToDropoff
                currentPosition = dropoff
            }
        }
        remainingLoads = removeLoads(remainingLoads, driverLoads)
        drivers = append(drivers, driverLoads)
    }
    return drivers
}