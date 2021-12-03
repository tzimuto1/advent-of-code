
package main

import (
    "fmt"
    "os"
)

func part1() {
    input, err := GetFileLines("resources/day2.txt")

    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }

    resultantForward := 0
    resultantDepth := 0
    for _, line := range input {
        distx := 0
        disty := 0
        motionType := line[0]
        if motionType == 'f' {
            _, err := fmt.Sscanf(line, "forward %d", &distx)
            ExitOnError(err)
        } else if (motionType == 'u') {
            _, err := fmt.Sscanf(line, "up %d", &disty)
            ExitOnError(err)
            disty = -1 * disty      
        } else if (motionType == 'd') {
            _, err := fmt.Sscanf(line, "down %d", &disty)
            if err != nil {
                fmt.Println(err)
                os.Exit(2)
            }
        }
        resultantForward += distx
        resultantDepth += disty
    }

    fmt.Printf("Answer: %d\n", resultantForward * resultantDepth)
    // Answer: 1654760
}

func part2() {
    input, err := GetFileLines("resources/day2.txt")

    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }

    resultantForward := 0
    resultantDepth := 0
    aim := 0
    for _, line := range input {
        x := 0
        motionType := line[0]
        if motionType == 'f' {
            _, err := fmt.Sscanf(line, "forward %d", &x)
            ExitOnError(err)
            resultantForward += x
            resultantDepth += (aim * x)
        } else if (motionType == 'u') {
            _, err := fmt.Sscanf(line, "up %d", &x)
            ExitOnError(err)
            aim -= x    
        } else if (motionType == 'd') {
            _, err := fmt.Sscanf(line, "down %d", &x)
            ExitOnError(err)
            aim += x
        }
    }

    fmt.Printf("Answer: %d\n", resultantForward * resultantDepth)
    // Answer: 1956047400
}

func main() {
    // part1()
    part2()
}