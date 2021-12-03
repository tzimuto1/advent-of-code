package main

import (
    "fmt"
    "log"
    "os"
    "strconv"
)

func getDepths() []int {
    fileLines, err := GetFileLines("resources/day1.txt")
    if err != nil {
        log.Fatal(err)
        os.Exit(2)
    }

    var depthVals []int
    for _, line := range fileLines {
        depth, err := strconv.Atoi(line)

        if err != nil {
            fmt.Println(err)
            os.Exit(2)
        }

        depthVals = append(depthVals, depth)

    }

    return depthVals
}

func part1() {
    depthVals, err := GetIntsFromFile("resources/day1.txt")

    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }

    // count number of times depth increases
    numTimesDepthIncrease := 0
    prevDepth := depthVals[0]
    for _, depth := range depthVals[1:] {
        if depth > prevDepth {
            numTimesDepthIncrease++;
        }
        prevDepth = depth;
    }

    fmt.Printf("Answer: %d\n", numTimesDepthIncrease)
    // Answer: 1298

}

func part2() {
    depths, err := GetIntsFromFile("resources/day1.txt")

    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }

    numTimesDepthSumIncrease := 0
    prevDepthSum := depths[0] + depths[1] + depths[2]
    for i := 3; i < len(depths); i++ {
        depthSum := prevDepthSum - depths[i-3] + depths[i]
        if depthSum > prevDepthSum {
            numTimesDepthSumIncrease++
        }
    }

    fmt.Printf("Answer: %d\n", numTimesDepthSumIncrease)
    // Answer: 1248
}

func main() {
    part1()
    part2()
}