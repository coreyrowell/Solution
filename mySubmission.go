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

func calculateDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
    d1 := x2 - x1
    d2 := y2 - y1
    distance := math.Sqrt((d1*d1) + (d2*d2))
    distanceRounded := roundToDecimalPlaces(distance, 2)
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

func assignLoadsToDriver(loads []Load) {

    // while we have remaining loads --
    //     assign a new driver who's location is the depot (0, 0)
    //     look at each load --
    //         calculate distance from drivers location to first pickup
    //         calculate distance from drivers pickup location to dropoff
    //         calculate distance from drivers dropoff location back to depot
    //         if the time to make the above round trip does not exceed 12 hours --
    //             (this load can be added to drivers schedule)
    //             add the load to drivers schedule
    //             track time to make the trip
    //             remove the load from our remaining loads
    //     all loads for the driver have been checked
    //     find a new driver if there are remaining loads
}

func main() {
    // @TODO - error check args, ensure filename arg is provided
    args := os.Args
    fileName := args[1]
    loads := readInput(fileName)
    assignLoadsToDriver(loads)
}