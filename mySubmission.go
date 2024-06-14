package main

import (
	"os"
	"solution/myCode/load"
	"solution/myCode/strategies"
	"solution/myCode/utils"
)

func main() {
    // @TODO - error check args, ensure filename arg is provided
    args := os.Args
    fileName := args[1]
    loads := utils.ReadInput(fileName)

    // Create the context with the desired strategy
    loadAssigner := &load.LoadAssigner{}
    // Use Farthest Pickup First strategy
    loadAssigner.SetStrategy(&strategies.FarthestPickupStrategy{})
    
    drivers := loadAssigner.AssignLoads(loads)
    utils.PrintExpectedOutput(drivers)
}