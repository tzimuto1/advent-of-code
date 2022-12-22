import string


PRIORITIES = {
    character: priority 
    for priority, character in enumerate(string.ascii_lowercase + string.ascii_uppercase, 1)
}

def part1(rucksacks_raw):
    # 7990
    rucksacks =  [(line[0:len(line) // 2], line[len(line) // 2:]) for line in rucksacks_raw]
    priority_sum = 0
    for rucksack in rucksacks:
        common_items = set(rucksack[0]).intersection(set(rucksack[1]))
        priority_sum += sum(PRIORITIES[item] for item in common_items)
    print(priority_sum)


def part2(rucksacks):
    # 2602
    groups = [rucksacks[i:i+3] for i in range(0, len(rucksacks), 3)]
    priority_sum = 0
    for group in groups:
        common_items = set(group[0]).intersection(set(group[1])).intersection(set(group[2]))
        priority_sum += sum(PRIORITIES[item] for item in common_items)
    print(priority_sum)

def main():
    with open("resources/day03.txt", "r") as f:
        rucksacks = f.read().strip().split("\n")
        part1(rucksacks)
        part2(rucksacks)

if __name__ == '__main__':
    main()