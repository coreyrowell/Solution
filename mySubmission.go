package main

import (
	"fmt"
	"math"
)

func roundToDecimalPlaces(num float64, decimalPlaces int) float64 {
    // https://gosamples.dev/round-float/
    ratio := math.Pow(10, float64(decimalPlaces))
    return math.Round(num*ratio) / ratio
}

func calculateDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
    d1 := x2 - x1
    d2 := y2 - y1
    distance := math.Sqrt((d1*d1) + (d2*d2))
    distanceRounded := roundToDecimalPlaces(distance, 2)
    return distanceRounded
}

func main() {
    fmt.Println("test")
}