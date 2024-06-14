package utils

import (
	"math"
)

func RoundToDecimalPlaces(num float64, decimalPlaces int) float64 {
    // https://gosamples.dev/round-float/
    ratio := math.Pow(10, float64(decimalPlaces))
    return math.Round(num*ratio) / ratio
}