import ../common/common
import std/strutils

type
    Tree = ref object
        height: int
        i: int
        j: int
        visible: bool
        lastViewLeft: int
        lastViewUp: int
        lastViewRight: int
        lastViewDown: int


proc getGrid(lines: var seq[string]): seq[seq[Tree]] =
    var grid = newSeq[seq[Tree]](len(lines))
    for j, line in lines:
        var row = newSeq[Tree](len(line))
        for i, treeHeightChar in line:
            var tree = Tree(height: parseInt($treeHeightChar), 
                            i: i,
                            j: j,
                            visible: false,
                            lastViewLeft: i,
                            lastViewUp: j,
                            lastViewRight: i,
                            lastViewDown: j)
            row[i] = tree
        grid[j] = row
    grid

proc part1(lines: var seq[string]) =
    # 1829
    var grid = getGrid(lines)

    # from left
    for j in 0 ..< len(grid):
        var prevHeight = low(int)
        for i in 0 ..< len(grid[j]):
            if grid[j][i].height > prevHeight:
                prevHeight = grid[j][i].height
                grid[j][i].visible = true
            elif grid[j][i].height == prevHeight:
                continue

    # from top
    for i in 0 ..< len(grid[0]):
        var prevHeight = low(int)
        for j in 0 ..< len(grid):
            if grid[j][i].height > prevHeight:
                prevHeight = grid[j][i].height
                grid[j][i].visible = true
            elif grid[j][i].height == prevHeight:
                continue

    # from right
    for j in 0 ..< len(grid):
        var prevHeight = low(int)
        for i in countdown(len(grid[j]) - 1, 0):
            if grid[j][i].height > prevHeight:
                prevHeight = grid[j][i].height
                grid[j][i].visible = true
            elif grid[j][i].height == prevHeight:
                continue

    # from bottom
    for i in 0 ..< len(grid[0]):
        var prevHeight = low(int)
        for j in countdown(len(grid) - 1, 0):
            if grid[j][i].height > prevHeight:
                prevHeight = grid[j][i].height
                grid[j][i].visible = true
            elif grid[j][i].height == prevHeight:
                continue

    # result
    var numVisibleTrees = 0
    for j in 0 ..< len(grid):
        for i in 0 ..< len(grid[j]):
            if grid[j][i].visible:
                numVisibleTrees += 1

    echo(numVisibleTrees)

proc part2(lines: var seq[string]) =
    # 291840
    var grid = getGrid(lines)

    # from left
    for j in 0 ..< len(grid):
        for i in 1 ..< len(grid[j]):
            var tree = grid[j][i]
            var otherTree = grid[j][i-1]
            while tree.height > otherTree.height and otherTree.i != 0:
                otherTree = grid[j][otherTree.lastViewLeft]

            tree.lastViewLeft = otherTree.i
            # grid[j][i] = tree

    # from top
    for i in 0 ..< len(grid[0]):
        for j in 1 ..< len(grid):
            var tree = grid[j][i]
            var otherTree = grid[j-1][i]
            while tree.height > otherTree.height and otherTree.j != 0:
                otherTree = grid[otherTree.lastViewUp][i]

            tree.lastViewUp = otherTree.j
            # grid[j][i] = tree
            

    # from right
    for j in 0 ..< len(grid):
        for i in countdown(len(grid[j]) - 2, 0):
            var tree = grid[j][i]
            var otherTree = grid[j][i+1]
            while tree.height > otherTree.height and otherTree.i != len(grid[j])-1:
                otherTree = grid[j][otherTree.lastViewRight]

            tree.lastViewRight = otherTree.i
            # grid[j][i] = tree

    # from bottom
    for i in 0 ..< len(grid[0]):
        for j in countdown(len(grid) - 2, 0):
            var tree = grid[j][i]
            var otherTree = grid[j+1][i]
            while tree.height > otherTree.height and otherTree.j != len(grid)-1:
                otherTree = grid[otherTree.lastViewDown][i]

            tree.lastViewDown = otherTree.j
            # grid[j][i] = tree

    # scenic scores
    var result = 0
    for j in 0 ..< len(grid):
        for i in 0 ..< len(grid[j]):
            var tree = grid[j][i]
            var scenicScore = (tree.i - tree.lastViewLeft) * (tree.j - tree.lastViewUp) * (tree.i - tree.lastViewRight) * (tree.j - tree.lastViewDown)
            if scenicScore > result:
                result = scenicScore
            

    echo(result)


proc main =
    var lines : seq[string]= getFile("resources/day08.txt")
    part1(lines)
    part2(lines)

main()