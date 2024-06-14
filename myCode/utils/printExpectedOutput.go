package utils

import "fmt"

func PrintExpectedOutput(drivers [][]int) {
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