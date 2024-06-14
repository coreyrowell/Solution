package main

import (
	"os"
)

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