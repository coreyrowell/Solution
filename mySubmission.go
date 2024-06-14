package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

func calculateDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
    d1 := x2 - x1
    d2 := y2 - y1
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

func assignLoadsToDriver(loads []Load) [][]int {
    drivers := [][]int{}
    remainingLoads := make([]Load, len(loads))
    copy(remainingLoads, loads)

    for len(remainingLoads) > 0 {
        driverLoads := []int{}
        currentTime := 0.0
        currentPosition := Point{X: 0, Y: 0}

        // Sort remainingLoads such that loads with endpoints further from current position come later in list
        sort.Slice(remainingLoads, func(i, j int) bool {
			distI := calculateDistance(currentPosition.X, currentPosition.Y, remainingLoads[i].End.X, remainingLoads[i].End.Y)
			distJ := calculateDistance(currentPosition.X, currentPosition.Y, remainingLoads[j].End.X, remainingLoads[j].End.Y)
			return distI > distJ
		})

        for _, load := range remainingLoads {
            loadNumber := load.Number
            pickup := load.Start
            dropoff := load.End
            timeToPickup := calculateDistance(currentPosition.X, currentPosition.Y, pickup.X, pickup.Y)
            timeToDropoff := calculateDistance(pickup.X, pickup.Y, dropoff.X, dropoff.Y)
            timeToDepot := calculateDistance(dropoff.X, dropoff.Y, 0, 0)

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

func main() {
    // @TODO - error check args, ensure filename arg is provided
    args := os.Args
    fileName := args[1]
    loads := readInput(fileName)
    drivers := assignLoadsToDriver(loads)
    printExpectedOutput(drivers)
}