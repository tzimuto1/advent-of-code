import ../common/common

import std/strutils

proc adjustSum(x: int, cycleCount: int): int =
    if cycleCount == 20 or (cycleCount - 20) mod 40 == 0 and cycleCount <= 220:
        return cycleCount * x
    else:
        return 0


proc part1(lines: var seq[string]) =
    var cycleCount = 1
    var x = 1
    var result = 0
    for line in lines:
        var tokens = line.split(" ")
        if tokens[0] == "noop":
            result += adjustSum(x, cycleCount)
            cycleCount += 1
        else:
            result += adjustSum(x, cycleCount)
            cycleCount += 1

            result += adjustSum(x, cycleCount)
            var increment = parseInt(tokens[1])
            x += increment
            cycleCount += 1

    echo(result)

type
    Screen = array[6, array[40, char]]

proc updateScreen(screen: var Screen, cycleCount: int, x: int) =
    let cycleIndex = cycleCount - 1 
    let spriteIndex = x - 1
    var row = (cycleIndex div 40) mod 6
    var col = cycleIndex mod 40

    if col-1 <= spriteIndex and spriteIndex <= col+1:
        screen[row][col] = '#'
    else:
        screen[row][col] = '.'

proc printScreen(screen: var Screen) =
    for row in 0 .. 5:
        for col in 0 .. 39:
            stdout.write(screen[row][col])
        stdout.write('\n')



proc part2(lines: var seq[string]) =
    # PLULKBZH
    var screen: Screen

    var cycleCount = 1
    var x = 1
    var result = 0
    for line in lines:
        var tokens = line.split(" ")
        if tokens[0] == "noop":
            updateScreen(screen, cycleCount, x)
            cycleCount += 1
        else:
            updateScreen(screen, cycleCount, x)
            cycleCount += 1

            var xIncrement = parseInt(tokens[1])
            x += xIncrement

            updateScreen(screen, cycleCount, x)
            cycleCount += 1
    printScreen(screen)

proc main =
    var lines : seq[string]= getFile("resources/day10.txt")
    part1(lines)
    part2(lines)

main()
