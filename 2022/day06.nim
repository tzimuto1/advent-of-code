import ../common/common

import std/tables


proc getCharacterCountAtFirstMarker(stream: var string, markerLength: int) : int =
    var counts = initTable[char, int]()

    var i = 0;

    while i < markerLength:
        if not counts.hasKey(stream[i]):
            counts[stream[i]] = 1
        else:
            counts[stream[i]] += 1
        i += 1

    var result : int
    while i < len(stream):
        var startOfOldWindow = stream[i-markerLength]
        counts[startOfOldWindow] -= 1
        if counts[startOfOldWindow] == 0:
            counts.del(startOfOldWindow)

        var endOfNewWindow = stream[i]
        if not counts.hasKey(endOfNewWindow):
            counts[endOfNewWindow] = 1
        else:
            counts[endOfNewWindow] += 1

        if len(counts) == markerLength:
            result = i + 1
            break
        i += 1

    result


proc part1(stream: var string) =
    # 1892
    echo(getCharacterCountAtFirstMarker(stream, 4))

proc part2(stream: var string) =
    # 2313
    echo(getCharacterCountAtFirstMarker(stream, 14))


proc main =
    var lines : seq[string]= getFile("resources/day06.txt")
    part1(lines[0])
    part2(lines[0])

main()