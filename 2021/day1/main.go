package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func getDepths() []int {
    // read the ints
    file, err := os.Open("input");
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var depthVals []int
    for scanner.Scan() {
        depthStr := scanner.Text()
        depth, err := strconv.Atoi(depthStr)

        if err != nil {
            fmt.Println(err)
            os.Exit(2)
        }

        depthVals = append(depthVals, depth)

    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
        os.Exit(2)
    }

    return depthVals
}

func part1() {
    depthVals := getDepths()

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
    depths := getDepths()

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
    // part1()
    part2()
}