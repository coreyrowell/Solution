package strategies

import (
	"solution/myCode/load"
	"solution/myCode/utils"
)

type UnsortedLoadsPickupStrategy struct{}

func (s *UnsortedLoadsPickupStrategy) AssignLoadsToDriver(loads []load.Load) [][]int {
    drivers := [][]int{}
    remainingLoads := make([]load.Load, len(loads))
    depot := load.Point{X:0, Y:0}
    copy(remainingLoads, loads)

    for len(remainingLoads) > 0 {
        driverLoads := []int{}
        currentTime := 0.0
        currentPosition := load.Point{X: 0, Y: 0}

        for _, load := range remainingLoads {
            loadNumber := load.Number
            pickup := load.Start
            dropoff := load.End
            timeToPickup := utils.CalculateDistance(currentPosition, pickup)
            timeToDropoff := utils.CalculateDistance(pickup, dropoff)
            timeToDepot := utils.CalculateDistance(dropoff, depot)

            if (currentTime + timeToPickup + timeToDropoff + timeToDepot) <= 720 {
                driverLoads = append(driverLoads, loadNumber)
                currentTime += timeToPickup + timeToDropoff
                currentPosition = dropoff
            }
        }
        remainingLoads = utils.RemoveLoads(remainingLoads, driverLoads)
        drivers = append(drivers, driverLoads)
    }
    return drivers
}