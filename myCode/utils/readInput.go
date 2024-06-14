package utils

import (
	"bufio"
	"log"
	"os"
	"solution/myCode/load"
)

func ReadInput(fileName string) []load.Load {
    // @TODO - return err if file not exists, or on parse error
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    loads := []load.Load{}
    isFirstLine := true
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        if isFirstLine {
            isFirstLine = false
            continue // Skip header
        }
        load := ParseLoadFromString(scanner.Text())
        loads = append(loads, load)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return loads
}