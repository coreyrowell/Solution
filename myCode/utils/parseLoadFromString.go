package utils

import (
	"fmt"
	"solution/myCode/load"
)

func ParseLoadFromString(line string) load.Load {
    // https://pkg.go.dev/fmt#example-Sscanf

    // Extract the load number
    var number int
	_, err := fmt.Sscanf(line, "%d", &number)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Extract the start and end points
	var startX, startY, endX, endY float64
	_, err = fmt.Sscanf(line, "%d (%f,%f) (%f,%f)",
		&number, &startX, &startY, &endX, &endY)
	if err != nil {
		fmt.Println("Error:", err)
	}

    load := load.Load{
		Number: number,
		Start: load.Point{
			X: startX,
			Y: startY,
		},
		End: load.Point{
			X: endX,
			Y: endY,
		},
	}

    return load
}