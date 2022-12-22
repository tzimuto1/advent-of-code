import ../common/common

import hashes
import std/sequtils
import std/sets
import std/strutils
import std/tables

type
    Point = object
        i: int
        j: int

type Vector = Point

proc hash(t: var Point): Hash =
    var h: Hash = 0
    h = h !& hash(t.i)
    h = h !& hash(t.j)
    result = !$h


proc `==`(this: var Point, that: var Point): bool =
    result = this.i == that.i and this.j == that.j


let DIRECTION_TO_VECTOR = {
    "L": Vector(i: -1, j: 0),
    "U": Vector(i: 0, j: 1),
    "R": Vector(i: 1, j: 0),
    "D": Vector(i: 0, j: -1)
}.toTable

proc distanceAdjustment(diff: int): int =
    if diff == 2:
        return 1
    elif diff == -2:
        return -1
    else:
        return diff

proc move(points: var seq[Point], 
          vector: Vector,
          distance: int,
          tails: var HashSet[Point]) =
    var d = 0
    while d < distance:
        if vector.i != 0:
            points[0].i = points[0].i + vector.i
        else:
            points[0].j = points[0].j + vector.j

        for i in 1 ..< len(points):

            var changeVector = Vector(
                i: points[i-1].i - points[i].i, 
                j: points[i-1].j - points[i].j
            )
            if abs(changeVector.i) == 2 or abs(changeVector.j) == 2:
                points[i].i += distanceAdjustment(changeVector.i)
                points[i].j += distanceAdjustment(changeVector.j)

                if i == len(points) - 1:
                    if points[i] notin tails:
                        tails.incl(points[i])

        d += 1

proc common(lines: var seq[string], numItems: int) =
    var points = repeat(Point(i: 0, j: 0), numItems)

    var tails = initHashSet[Point]()
    tails.incl(points[0])

    var i = 0
    for line in lines:
        var tokens = line.split(" ")
        var direction = tokens[0]
        var distance = parseInt(tokens[1])
        move(points, DIRECTION_TO_VECTOR[direction], distance, tails)

    echo(len(tails))


proc part1(lines: var seq[string]) =
    # 6087
    common(lines, 2)

proc part2(lines: var seq[string]) =
    # 2493
    common(lines, 10)


proc main =
    var lines : seq[string]= getFile("resources/day09.txt")
    part1(lines)
    part2(lines)

main()
