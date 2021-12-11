package main

import (
	"fmt"
	"math"
)


type Point struct {
	x, y int
}


func getNumPointsWithAtLeastTwoLines(nonDiagonalOnly bool) int {
	lines, err := GetFileLines("resources/day5.txt")
	ExitOnError(err)

	pointToLineCount := make(map[Point]int)

	for _, lineStr := range lines {
		var x1, y1, x2, y2 int
		_, err := fmt.Sscanf(lineStr, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		ExitOnError(err)

		if nonDiagonalOnly && x1 != x2 && y1 != y2 {
			continue
		}

		dx := 0
		dy := 0
		steps := 0

		if nonDiagonalOnly {
			if x1 == x2 {
				steps = int(math.Abs(float64(y2 - y1)))
				dx = 0
				if y2 > y1 {
					dy = 1
				} else {
					dy = -1
				}
			} else if y1 == y2 {
				steps = int(math.Abs(float64(x2 - x1)))
				dy = 0
				if x2 > x1 {
					dx = 1
				} else {
					dx = -1
				}
			}
		} else {
			steps = int(math.Max(math.Abs(float64(y2 - y1)), math.Abs(float64(x2 - x1))))
			if y2 > y1 {
				dy = 1
			} else if y2 < y1 {
				dy = -1
			}
			if x2 > x1 {
				dx = 1
			} else if x2 < x1 {
				dx = -1
			}
		}

		newX := x1
		newY := y1
		for i := 0; i < steps + 1; i++ {
			newPoint := Point{newX, newY}
			//fmt.Println(lineStr, newPoint)
			pointToLineCount[newPoint] += 1
			newX += dx
			newY += dy
		}
		//break
	}

	//fmt.Println(pointToLineCount)

	numPointsWithAtLeastTwoLines := 0
	for _, lineCount := range pointToLineCount {
		if lineCount >= 2 {
			numPointsWithAtLeastTwoLines += 1
		}
	}

	return numPointsWithAtLeastTwoLines
}

func part1() {
	numPointsWithAtLeastTwoLines := getNumPointsWithAtLeastTwoLines(true)
	fmt.Printf("Answer: %d\n", numPointsWithAtLeastTwoLines)
	// 8060
}

func part2() {
	numPointsWithAtLeastTwoLines := getNumPointsWithAtLeastTwoLines(false)
	fmt.Printf("Answer: %d\n", numPointsWithAtLeastTwoLines)
	// 12999
}

func main() {
	part1()
	part2()
}
