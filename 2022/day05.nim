import std/tables
import std/strutils
import std/strscans
import std/deques

proc getFile : seq[string] =
    var lines : seq[string]
    let f = open("resources/day05.txt")
    defer:
        f.close()

    var line : string
    while f.read_line(line):
        lines.add(line)

    lines


proc part1(lines: seq[string]) =
    # JDTMRWCQJ
    var stacks = initTable[int, Deque[char]]()
    var instructions = false
    for line in lines:
        if not instructions:
            if line != "":
                if not line[1].isDigit():
                    var i = 0
                    var stackIndex = 1
                    while i < len(line):
                        if line[i] == '[':
                            if not stacks.hasKey(stackIndex):
                                stacks[stackIndex] = initDeque[char]()
                            stacks[stackIndex].addFirst(line[i + 1])
                        i += 4
                        stackIndex += 1
            else:
                instructions = true
        else:
            var count, fromStack, toStack : int
            if scanf(line, "move $i from $i to $i", count, fromStack, toStack):
                var r : int
                while r < count:
                    stacks[toStack].addLast(stacks[fromStack].popLast())
                    r += 1

    var result = newStringOfCap(9)
    for stackIndex in 1..9:
        add(result, stacks[stackIndex][^1])
    echo result


proc part2(lines: seq[string]) =
    # VHJDDCWRD
    var stacks = initTable[int, Deque[char]]()
    var instructions = false
    for line in lines:
        if not instructions:
            if line != "":
                if not line[1].isDigit():
                    var i = 0
                    var stackIndex = 1
                    while i < len(line):
                        if line[i] == '[':
                            if not stacks.hasKey(stackIndex):
                                stacks[stackIndex] = initDeque[char]()
                            stacks[stackIndex].addFirst(line[i + 1])
                        i += 4
                        stackIndex += 1
            else:
                instructions = true
        else:
            var count, fromStack, toStack : int
            if scanf(line, "move $i from $i to $i", count, fromStack, toStack):
                var cratesToRemove : Deque[char]
                var r : int
                while r < count:
                    cratesToRemove.addLast(stacks[fromStack].popLast())
                    r += 1
                r = 0
                while r < count:
                    stacks[toStack].addLast(cratesToRemove.popLast())
                    r += 1

    var result = newStringOfCap(9)
    for stackIndex in 1..9:
        add(result, stacks[stackIndex][^1])
    echo result

proc main =
    var lines = getFile()
    part1(lines)
    part2(lines)

main()



