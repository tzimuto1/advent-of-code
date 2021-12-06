
package main

import (
    "fmt"
    "regexp"
    "strconv"
    "strings"
)


type BoardCell struct {
    row, col int
    isMarked bool
}

func findBoardScore(board [5][]string, numbers []string) (int, int) {
    numberToBoardCell := make(map[string]*BoardCell)
    for row, rowNumbers := range board {
        for col, number := range rowNumbers {
            numberToBoardCell[number] = &BoardCell{row, col, false}
        }
    }

    var rowStatuses [5]int
    var colStatuses [5]int
    for numberIndex, number := range numbers {
        boardCell, isPresent := numberToBoardCell[number]
        if isPresent {
            boardCell.isMarked = true
            rowStatuses[boardCell.row]++
            colStatuses[boardCell.col]++

            if rowStatuses[boardCell.row] == 5 || colStatuses[boardCell.col] == 5 {
                sumUnmarkedNumbers := 0
                for num, cell := range numberToBoardCell {
                    if !cell.isMarked {
                        unMarkedNumber, _ := strconv.Atoi(num)
                        sumUnmarkedNumbers += unMarkedNumber
                    }
                }
                numberInt, _ := strconv.Atoi(number)
                return sumUnmarkedNumbers * numberInt, numberIndex
            }
        }
    }

    return -1, -1
}

func part1() {
	inputs, err := GetFileLines("resources/day4.txt")
    ExitOnError(err)

    inputsPtr := 0

    numbers := strings.Split(inputs[inputsPtr], ",")
    // skip space and move to boards
    inputsPtr += 1 + 1


    bestBoardScore := 0
    bestBoardIndex := len(numbers)
    for inputsPtr < len(inputs) {
        var board [5][]string
        for i := 0; i < 5; i++ {
            boardRow := regexp.MustCompile(`[0-9]+`).FindAllString(inputs[inputsPtr+i], -1)
            board[i] = boardRow
        }

        boardScore, boardIndex := findBoardScore(board, numbers)
        if boardIndex < bestBoardIndex {
            bestBoardIndex = boardIndex
            bestBoardScore = boardScore
        }

        // skip board and empty line
        inputsPtr += 5 + 1
    }

    fmt.Printf("Answer: %d\n", bestBoardScore)
    // 32844
}

func part2() {
    inputs, err := GetFileLines("resources/day4.txt")
    ExitOnError(err)

    inputsPtr := 0

    numbers := strings.Split(inputs[inputsPtr], ",")
    // skip space and move to boards
    inputsPtr += 1 + 1


    bestBoardScore := 0
    bestBoardIndex := -1
    for inputsPtr < len(inputs) {
        var board [5][]string
        for i := 0; i < 5; i++ {
            boardRow := regexp.MustCompile(`[0-9]+`).FindAllString(inputs[inputsPtr+i], -1)
            board[i] = boardRow
        }

        boardScore, boardIndex := findBoardScore(board, numbers)
        if boardIndex > bestBoardIndex {
            bestBoardIndex = boardIndex
            bestBoardScore = boardScore
        }

        // skip board and empty line
        inputsPtr += 5 + 1
    }

    fmt.Printf("Answer: %d\n", bestBoardScore)
    // 32844
}

func main() {
    // part1()
    part2()
}