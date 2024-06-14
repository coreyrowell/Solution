package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type Point struct {
    X float64
    Y float64
}

type Load struct {
    Number int
    Start Point
    End Point
}


func roundToDecimalPlaces(num float64, decimalPlaces int) float64 {
    // https://gosamples.dev/round-float/
    ratio := math.Pow(10, float64(decimalPlaces))
    return math.Round(num*ratio) / ratio
}

func calculateDistance(pointA Point, pointB Point) float64 {
    d1 := pointA.X - pointB.X
    d2 := pointA.Y - pointB.Y
    distance := math.Sqrt((d1*d1) + (d2*d2))
    distanceRounded := roundToDecimalPlaces(distance, 3)
    return distanceRounded
}

func parseLoadFromString(line string) Load {
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

    load := Load{
		Number: number,
		Start: Point{
			X: startX,
			Y: startY,
		},
		End: Point{
			X: endX,
			Y: endY,
		},
	}

    return load
}

func readInput(fileName string) []Load {
    // @TODO - return err if file not exists, or on parse error
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    loads := []Load{}
    isFirstLine := true
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        if isFirstLine {
            isFirstLine = false
            continue // Skip header
        }
        load := parseLoadFromString(scanner.Text())
        loads = append(loads, load)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return loads
}

func contains(numbers []int, number int) bool {
	for _, n := range numbers {
		if n == number {
			return true
		}
	}
	return false
}

func removeLoads(loads []Load, toRemove []int) []Load {
	result := []Load{}
	for _, load := range loads {
		if !contains(toRemove, load.Number) {
			result = append(result, load)
		}
	}
	return result
}

func printExpectedOutput(drivers [][]int) {
    for _,driverLoads := range drivers {
        fmt.Print("[")
        for i,loadNum := range driverLoads {
            fmt.Print(loadNum)
            if i < len(driverLoads)-1 {
                fmt.Print(",")
            }
        }
        fmt.Print("]")
        fmt.Print("\n")
    }
}

func main() {
    // @TODO - error check args, ensure filename arg is provided
    args := os.Args
    fileName := args[1]
    loads := readInput(fileName)

    // Create the context with the desired strategy
    loadAssigner := &LoadAssigner{}
    // Use Farthest Pickup First strategy
    loadAssigner.SetStrategy(&FarthestPickupStrategy{})
    
    drivers := loadAssigner.AssignLoads(loads)
    printExpectedOutput(drivers)
}