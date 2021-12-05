
package main

import (
    "fmt"
    "strconv"
)


func part1() {
	numbers, err := GetFileLines("resources/day3.txt")
    ExitOnError(err)

    var numOnes [12]int;
    for _, binaryNumber := range numbers {
    	for i, bit := range binaryNumber {
    		if bit == '1' {
    			numOnes[i]++
    		}
    	}
    }

    var halfNumInputs = len(numbers) / 2
    var gammaStr string;
    for _, count := range numOnes {
    	if count > halfNumInputs {
    		gammaStr += "1"
    	} else {
    		gammaStr += "0"
    	}
    }

    gamma, err := strconv.ParseInt(gammaStr, 2, 64)
    epsilon := (^gamma) & 0b111111111111

    fmt.Printf("Answer: %d\n", gamma * epsilon)
    // 775304
}

func getBitRows(numbers []string, bitColumn int, prevBitRows []int) ([]int, []int) {
	var bitOneRows []int
	var bitZeroRows []int

	for _, index := range prevBitRows {
		bit := numbers[index][bitColumn]
		if bit == '1' {
			bitOneRows = append(bitOneRows, index)
		} else {
			bitZeroRows = append(bitZeroRows, index)
		}
	}

	return bitOneRows, bitZeroRows
}

func part2() {
	numbers, err := GetFileLines("resources/day3.txt")
    ExitOnError(err)

    var bitRows[]int;
    for i := 0; i < len(numbers); i++ {
    	bitRows = append(bitRows, i)
    }

    var o2_generator_rating int64
    for i := 0; i < 12; i++ {
    	bitOneRows, bitZeroRows := getBitRows(numbers, i, bitRows)
    	if len(bitOneRows) >= len(bitZeroRows) {
    		bitRows = bitOneRows
    	} else {
    		bitRows = bitZeroRows
    	}

    	if len(bitRows) == 1 {
    		o2_generator_rating, err = strconv.ParseInt(numbers[bitRows[0]], 2, 64)
    		break
    	}
    }

    for i := 0; i < len(numbers); i++ {
    	bitRows = append(bitRows, i)
    }

    var co2_generator_rating int64
    for i := 0; i < 12; i++ {
    	bitOneRows, bitZeroRows := getBitRows(numbers, i, bitRows)
    	if len(bitOneRows) < len(bitZeroRows) {
    		bitRows = bitOneRows
    	} else {
    		bitRows = bitZeroRows
    	}

    	if len(bitRows) == 1 {
    		co2_generator_rating, err = strconv.ParseInt(numbers[bitRows[0]], 2, 64)
    		break
    	}
    }

    fmt.Printf("Answer: %d\n", o2_generator_rating * co2_generator_rating)
}


func main() {
    part1()
    part2()
}