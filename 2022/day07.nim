import ../common/common
import std/strutils
import std/strscans

type
    DirNode = tuple
        parentIndex: int
        totalSize: int


proc part1(lines: seq[string]) =
    # 1315285
    var sizes : seq[int]

    var result : int = 0;
    var i = 0
    while i < len(lines):
        var tokens = lines[i].split(' ')
        if tokens[0] == "$":
            var command = tokens[1]
            if command == "cd":
                var directory = tokens[2]
                if directory != "..":
                    sizes.add(0)
                else:
                    var childDirSize = sizes.pop()
                    if childDirSize <= 100_000:
                        result += childDirSize
                    sizes[^1] += childDirSize
            elif command == "ls":
                discard
        elif tokens[0][0].isDigit():
            var fileSize: int
            if scanf(lines[i], "$i", fileSize):
                sizes[^1] += fileSize
        elif tokens[0] == "dir":
            discard

        i += 1
    if sizes[0] <= 100_000:
        result += sizes[0]

    echo(result)


proc part2(lines: seq[string]) =
    # 9847279
    var sizes : seq[int]
    var allSizes : seq[int]

    var i = 0
    while i < len(lines):
        var tokens = lines[i].split(' ')
        if tokens[0] == "$":
            var command = tokens[1]
            if command == "cd":
                var directory = tokens[2]
                if directory != "..":
                    sizes.add(0)
                else:
                    var childDirSize = sizes.pop()
                    allSizes.add(childDirSize)
                    sizes[^1] += childDirSize
            elif command == "ls":
                discard
        elif tokens[0][0].isDigit():
            var fileSize: int
            if scanf(lines[i], "$i", fileSize):
                sizes[^1] += fileSize
        elif tokens[0] == "dir":
            discard

        i += 1

    while len(sizes) > 1:
        var childDirSize = sizes.pop()
        allSizes.add(childDirSize)
        sizes[^1] += childDirSize

    let MAX_CAPACITY = 70_000_000
    let SPACE_NEEDED = 30_000_000

    allSizes.add(sizes[0])


    let usedSpace = sizes[0]
    var availableSpace = MAX_CAPACITY - sizes[0]
    var minimumNeededSpace = SPACE_NEEDED - availableSpace

    var result = high(int)
    for size in allSizes:
        if size >= minimumNeededSpace and size < result:
            result = size
    echo(result)

proc main =
    var lines : seq[string]= getFile("resources/day07.txt")
    part1(lines)
    part2(lines)

main()