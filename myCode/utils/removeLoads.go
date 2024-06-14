package utils

import (
	"solution/myCode/load"
)

func RemoveLoads(loads []load.Load, toRemove []int) []load.Load {
	result := []load.Load{}
	for _, load := range loads {
		if !contains(toRemove, load.Number) {
			result = append(result, load)
		}
	}
	return result
}