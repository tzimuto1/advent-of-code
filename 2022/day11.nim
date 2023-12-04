import ../common/common

import std/deques
import std/strutils

type
    Test = object
        divisor: int
        trueMonkey: int
        falseMonkey: int

type
    Monkey = object
        id: int
        items: Deque[int]
        operation: seq[string]
        test: Test
        numInspections: int

proc getNextInstruction(lines: var seq[string], index: int): string =
    return lines[index].strip(trailing = false) 


proc getMonkeys(lines: var seq[string]): seq[Monkey] =
    var monkeys: seq[Monkey]
    var i = 0
    while i < len(lines):
        var instruction = getNextInstruction(lines, i)

        # monkey ID
        var monkeyTokens = instruction.split(" ")
        var id = parseInt(monkeyTokens[1][0..^2])
        i += 1

        # monkey items
        var items = initDeque[int]()
        instruction = getNextInstruction(lines, i)
        var itemsStr = instruction.split(": ")[1].split(", ")
        for itemStr in itemsStr:
            items.addLast(itemStr.parseInt())
        i += 1

        # monkey operation
        instruction = getNextInstruction(lines, i)
        var operation = instruction.split(": new = ")[1].split(" ")
        i += 1

        # monkey test
        instruction = getNextInstruction(lines, i)
        var divisor = instruction.split(" ")[3].parseInt()
        i += 1

        instruction = getNextInstruction(lines, i)
        var trueMonkey = instruction.split(" ")[5].parseInt()
        i += 1

        instruction = getNextInstruction(lines, i)
        var falseMonkey = instruction.split(" ")[5].parseInt()
        i += 1

        var test = Test(divisor: divisor, trueMonkey: trueMonkey, falseMonkey: falseMonkey)

        # monkey separator
        i += 1

        var monkey: Monkey = Monkey(id: id, items: items, operation: operation, test: test,
                                    numInspections: 0)
        monkeys.add(monkey)

    monkeys


proc getOperationValue(monkey: Monkey, item: int): int =
    let oldValue = item
    var newValue = oldValue

    var rightOperand: int
    if monkey.operation[2] == "old":
        rightOperand = oldValue
    else:
        rightOperand = monkey.operation[2].parseInt()

    case monkey.operation[1]:
        of "+":
            newValue += rightOperand
        of "*":
            newValue *= rightOperand

    return newValue


proc getTestResult(monkey: Monkey, newValue: int): bool =
    return newValue mod monkey.test.divisor == 0

proc part1(lines: var seq[string]) = 
    var monkeys = getMonkeys(lines)

    for round in 1..20:
        for monkeyId, monkey in monkeys:
            while len(monkeys[monkeyId].items) > 0:
                var item = monkeys[monkeyId].items.popFirst()
                var newItem = getOperationValue(monkey, item)

                # relief adjustment
                newItem = newItem div 3

                # determine receiving monkey and update
                var testResult = getTestResult(monkey, newItem)
                var receivingMonkeyId: int
                if testResult:
                    receivingMonkeyId = monkey.test.trueMonkey
                else:
                    receivingMonkeyId = monkey.test.falseMonkey

                monkeys[receivingMonkeyId].items.addLast(newItem)
                monkeys[monkeyId].numInspections += 1

    var maxNumInspections1 = -1
    var maxNumInspections2 = -1
    for monkey in monkeys:
        if monkey.numInspections >= maxNumInspections1:
            maxNumInspections2 = maxNumInspections1
            maxNumInspections1 = monkey.numInspections
        elif monkey.numInspections > maxNumInspections2:
            maxNumInspections2 = monkey.numInspections

    echo(maxNumInspections1 * maxNumInspections2)

proc part2(lines: var seq[string]) =
    var monkeys = getMonkeys(lines)

    for round in 1..10_000:
        for monkeyId, monkey in monkeys:
            while len(monkeys[monkeyId].items) > 0:
                var item = monkeys[monkeyId].items.popFirst()
                var newItem = getOperationValue(monkey, item)

                # relief adjustment
                newItem = newItem div high(int)

                # determine receiving monkey and update
                var testResult = getTestResult(monkey, newItem)
                var receivingMonkeyId: int
                if testResult:
                    receivingMonkeyId = monkey.test.trueMonkey
                else:
                    receivingMonkeyId = monkey.test.falseMonkey

                monkeys[receivingMonkeyId].items.addLast(newItem)
                monkeys[monkeyId].numInspections += 1
    
    var maxNumInspections1 = -1
    var maxNumInspections2 = -1
    for monkey in monkeys:
        if monkey.numInspections >= maxNumInspections1:
            maxNumInspections2 = maxNumInspections1
            maxNumInspections1 = monkey.numInspections
        elif monkey.numInspections > maxNumInspections2:
            maxNumInspections2 = monkey.numInspections

    echo(maxNumInspections1, " ", maxNumInspections2)
    echo(maxNumInspections1 * maxNumInspections2)
    echo(monkeys)

proc main =
    var lines : seq[string]= getFile("resources/day11.txt.test")
    part1(lines)
    part2(lines)

main()
