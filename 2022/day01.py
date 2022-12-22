import functools

def group_calories_by_elf(all_calories_str):
    all_elf_calories = []
    elf_calories = []
    for elf_calories_str in all_calories_str:
        if elf_calories_str != '':
            elf_calories.append(int(elf_calories_str))
        else:
            all_elf_calories.append(elf_calories)
            elf_calories = []
    return all_elf_calories


def part1(all_elf_calories):
    max_calories = functools.reduce(
        lambda elf_total1, elf_total2: max(elf_total1, elf_total2),
        map(lambda elf_calories: sum(elf_calories), all_elf_calories)
    )
    print(max_calories)

def part2(all_elf_calories):
    def get_min_index(_array):
        min_index = 0
        min_value = _array[0]
        for i in range(1, len(_array)):
            if _array[i] < min_value:
                min_value = _array[i]
                min_index = i
        return min_index

    top_3_totals = [sum(all_elf_calories[0]), sum(all_elf_calories[1]), sum(all_elf_calories[2])]
    for i in range(3, len(all_elf_calories)):
        min_index = get_min_index(top_3_totals)
        elf_total = sum(all_elf_calories[i])
        if elf_total > top_3_totals[min_index]:
            top_3_totals[min_index] = elf_total
    print(sum(top_3_totals))



def main():
    with open("resources/day01.txt", "r") as f:
        all_calories_str = f.read().split("\n")
        all_elf_calories = group_calories_by_elf(all_calories_str)
        # part1(all_elf_calories)
        part2(all_elf_calories)
        

if __name__ == '__main__':
    main()