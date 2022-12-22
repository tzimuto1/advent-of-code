package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func part1() {
	input, err := GetFileLines("resources/day7.txt")
	ExitOnError(err)

	re := regexp.MustCompile(",")

	crabPositionsStr := re.Split(input[0], -1)
	numCrabs := len(crabPositionsStr)

	crabPosition2count := make(map[int]int)
	minCrabPosition := math.MaxInt64
	maxCrabPosition := 0
	maxFuelCost := 0

	for _, positionStr := range crabPositionsStr {
		position, err := strconv.Atoi(positionStr)
		ExitOnError(err)
		crabPosition2count[position] += 1
		minCrabPosition = Min(minCrabPosition, position)
		maxCrabPosition = Max(maxCrabPosition, position)
		maxFuelCost += position
	}
	fmt.Printf("Mix crab position %d, max crab position %d, max fuel cost %d\n", 
		       minCrabPosition, maxCrabPosition, maxFuelCost)

	minFuelCost := maxFuelCost
	fuelCost := maxFuelCost
	numLeftwardCrabs := 0
	numCurrPosCrabs := crabPosition2count[0]
	numRightwardCrabs := numCrabs - numCurrPosCrabs

	if (numLeftwardCrabs + numCurrPosCrabs + numRightwardCrabs != numCrabs) {
		panic("The crab counts don't match. Initial")
	}

	for position := Max(minCrabPosition, 1); position <= maxCrabPosition; position++{
		count, isPresent := crabPosition2count[position]
		if isPresent {
			numLeftwardCrabs = numLeftwardCrabs + numCurrPosCrabs
			numCurrPosCrabs = count
			numRightwardCrabs = numRightwardCrabs - count
		} else {
			numLeftwardCrabs = numLeftwardCrabs + numCurrPosCrabs
			numCurrPosCrabs = 0
		}
		fuelCost = fuelCost + numLeftwardCrabs - numRightwardCrabs - numCurrPosCrabs
		minFuelCost = Min(minFuelCost, fuelCost)

		if (numLeftwardCrabs + numCurrPosCrabs + numRightwardCrabs != numCrabs) {
			panic("The crab counts don't match")
		}
	}
	fmt.Printf("Answer: %d\n", minFuelCost)
}

func main() {
	part1()
}
