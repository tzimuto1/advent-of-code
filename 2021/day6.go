package main

import (
	"fmt"
	"regexp"
	"strconv"
	//"strings"
)


type FishInfo struct {
	originalFishRemainingDays int 
	numOriginalFish int
	generations[] int
}

func part1() {
	input, err := GetFileLines("resources/day6.txt")
	ExitOnError(err)

	re := regexp.MustCompile(",")

	fishRemainingDaysList := re.Split(input[0], -1)

	originalRemainingDaysToFishCount := make(map[int]int)
	for _, daysStr := range fishRemainingDaysList {
		days, _ := strconv.Atoi(daysStr)
		originalRemainingDaysToFishCount[days] += 1
	}

	var fishInfo []FishInfo
	for days, fishCount := range originalRemainingDaysToFishCount {
		fishInfo = append(fishInfo, FishInfo{days, fishCount, []int{days}})
	}

	for d := 0; d < 80; d++ {
		for i := 0; i < len(fishInfo); i++ {
			numGenerations := len(fishInfo[i].generations)
			for j := 0; j < numGenerations; j++ {
				if fishInfo[i].generations[j] == 0 {
					fishInfo[i].generations = append(fishInfo[i].generations, 8)
					fishInfo[i].generations[j] = 6
				} else {
					fishInfo[i].generations[j] -= 1
				}
			}
		}
	}

	totalFish := 0
	for i := 0; i < len(fishInfo); i++ {
		totalFish += fishInfo[i].numOriginalFish * len(fishInfo[i].generations)
	}

	fmt.Printf("Answer: %d\n", totalFish)
	// 391671
	
}


func f(timeFishCreated int, numDays int, memo map[int]int) int {
	fishCount, isPresent := memo[timeFishCreated]
	if isPresent {
		return fishCount
	}

	if timeFishCreated > numDays {
		return 0
	}

	newFishCount := (numDays + 8 - timeFishCreated) / 7 + f(timeFishCreated + 9, numDays, memo)

	for k := 0; k < ((numDays -8 - timeFishCreated) / 7); k++ {
		newFishCount += f(timeFishCreated + 7 * k + 16, numDays, memo)
	}

	memo[timeFishCreated] = newFishCount
	return newFishCount
}

func part1_modified() {
	input, err := GetFileLines("resources/day6.txt")
	ExitOnError(err)

	re := regexp.MustCompile(",")

	fishRemainingDaysList := re.Split(input[0], -1)

	originalRemainingDaysToFishCount := make(map[int]int)
	for _, daysStr := range fishRemainingDaysList {
		days, _ := strconv.Atoi(daysStr)
		originalRemainingDaysToFishCount[days] += 1
	}

	var fishInfo []FishInfo
	for days, fishCount := range originalRemainingDaysToFishCount {
		fishInfo = append(fishInfo, FishInfo{days, fishCount, []int{days}})
	}

	totalFishCount := 0
	for i := 0; i < len(fishInfo); i++ {
		memo := make(map[int]int)
		newFishCount := f(fishInfo[i].originalFishRemainingDays, 80 - fishInfo[i].originalFishRemainingDays, memo)
		totalFishCount += (newFishCount * fishInfo[i].numOriginalFish) + fishInfo[i].numOriginalFish
	} 

	fmt.Printf("Answer: %d\n", totalFishCount)
	// 391671
	// 438256

}

func main() {
	part1_modified()
}