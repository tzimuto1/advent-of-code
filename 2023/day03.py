
import string

import reader


DIGITS = set(string.digits)

def is_start_of_number(schematic, r, c):
	return schematic[r][c] in DIGITS


def extract_number(schematic, r, c):
	number = ''
	while c < len(schematic[0]) and is_start_of_number(schematic, r, c):
		number += schematic[r][c]
		c += 1
	return number


def is_digit_near_symbol(schematic, digit, r, c):
	R, C = len(schematic), len(schematic[0])
	for dr in (-1, 0, 1):
		for dc in (-1, 0, 1):
			r2, c2 = r + dr, c + dc
			if (0 <= r2 < R) and (0 <= c2 < C):
				neighbor = schematic[r2][c2]
				if neighbor not in DIGITS and neighbor != '.':
					return True
	return False


def is_part_number(schematic, number, r, c):
	for i in range(len(number)):
		if is_digit_near_symbol(schematic, number[i], r, c + i):
			return True
	return False


def part1(schematic):
	sum_of_part_numbers = 0
	for r in range(len(schematic)):
		c = 0
		while c < len(schematic[0]):
			if is_start_of_number(schematic, r, c):
				number = extract_number(schematic, r, c)
				if is_part_number(schematic, number, r, c):
					sum_of_part_numbers += int(number)
				c += len(number)
			else:
				c += 1
	# 525911
	print(sum_of_part_numbers)

def part2(schematic):
# 	schematic = reader.read_grid(
# 		"""
# 467..114..
# ...*......
# ..35..633.
# ......#...
# 617*......
# .....+.58.
# ..592.....
# ......755.
# ...$.*....
# .664.598.."""[1:], False)
# 	print(schematic)

	sum_of_gear_ratios = 0
	R, C = len(schematic), len(schematic[0])
	for r in range(R):
		for c in range(C):
			if schematic[r][c] != '*':
				continue
			visited_digit_neighbors = set()
			adjacent_numbers = []
			for dr in (-1, 0, 1):
				for dc in (-1, 0, 1):
					r2, c2 = r + dr, c + dc
					if ((0 <= r2 < R) 
						and (0 <= c2 < C) 
						and not (dr == 0 and dc == 0) 
						and (r2, c2) not in visited_digit_neighbors):
						if schematic[r2][c2] in DIGITS:
							c_num_start = c2
							visited_digit_neighbors.add((r2, c_num_start))
							while c_num_start > 0 and schematic[r2][c_num_start-1] in DIGITS:
								c_num_start -= 1
								visited_digit_neighbors.add((r2, c_num_start))
							c_num_end = c2
							while c_num_end < R - 1 and schematic[r2][c_num_end+1] in DIGITS:
								c_num_end += 1
								visited_digit_neighbors.add((r2, c_num_end))
							adjacent_numbers.append(int("".join(schematic[r2][c_num_start:c_num_end+1])))
			if len(adjacent_numbers) == 2:
				gear_ratio = adjacent_numbers[0] * adjacent_numbers[1]
				sum_of_gear_ratios += gear_ratio
	# 75805607
	print(sum_of_gear_ratios)



def main():
    schematic = reader.read_grid("resources/day03.txt")
    part1(schematic)
    part2(schematic)        


if __name__ == '__main__':
    main()
