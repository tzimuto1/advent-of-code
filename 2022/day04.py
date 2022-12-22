

def does_range_contain_range(a, b):
    return (a[0] <= b[0] and a[1] >= b[1])

def does_range_contain_point(r, point):
    return (r[0] <= point <= r[1]) or (r[0] == point) or (r[1] == point)

def part1(pairs):
    # 444
    num_fully_contained = 0
    for pair in pairs:
        (a, b) = pair
        if does_range_contain_range(a, b) or does_range_contain_range(b, a):
            num_fully_contained += 1 
    print(num_fully_contained)


def part2(pairs):
    # 686
    num_contained = 0
    for pair in pairs:
        (a, b) = pair
        if b[0] <= a[0] <= b[1] or a[0] in b or a[0] <= b[0] <= a[1] or b[0] in a:
        # if does_range_contain_point(a, b[0]) or does_range_contain_point(a, b[1]):
            print(a, b)
            num_contained += 1
    print(num_contained)
a[0] <= b[0] <= a[1] or a[0] == b[0] or a[1] == b[0]
#     ...........
#     #
    

def main():
    with open("resources/day04.txt", "r") as f:
        pairs = [
            tuple(
                tuple(
                    int(num) 
                    for num in elf.split("-")
                )
                for elf in pair.split(",")
            )
            for pair in f.read().strip().split("\n")
        ]
        part1(pairs)
        part2(pairs)

if __name__ == '__main__':
    main()
