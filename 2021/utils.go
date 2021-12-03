package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func GetFileLines(fileName string) ([]string, error) {
    file, err := os.Open(fileName);
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var fileLines []string
    for scanner.Scan() {
        line := scanner.Text()
        fileLines = append(fileLines, line)
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return fileLines, err
}

func GetIntsFromFile(fileName string) ([]int, error) {
    fileLines, err := GetFileLines(fileName)
    if err != nil {
        return nil, err
    }

    var values []int
    for _, line := range fileLines {
        value, err := strconv.Atoi(line)

        if err != nil {
            return nil, err
        }

        values = append(values, value)
    }

    return values, nil;
}

func ExitOnError(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }
}
