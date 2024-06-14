package utils

import (
	"math"
	"solution/myCode/load"
)

func CalculateDistance(pointA load.Point, pointB load.Point) float64 {
    d1 := pointA.X - pointB.X
    d2 := pointA.Y - pointB.Y
    distance := math.Sqrt((d1*d1) + (d2*d2))
    distanceRounded := RoundToDecimalPlaces(distance, 3)
    return distanceRounded
}