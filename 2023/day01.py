
import string


DIGITS = set(string.digits)

def part1(calibration_value_strings):
    calibration_values_sum = 0

    for calibration_value_string in calibration_value_strings:
        calibration_digits = [c for c in calibration_value_string if c in DIGITS]
        calibration_value = calibration_digits[0] + calibration_digits[-1]
        calibration_values_sum += int(calibration_value)
    print(calibration_values_sum)


def part2(calibration_value_strings):
    calibration_values_sum = 0

    replacements = {
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9"
    }

    for calibration_value_string in calibration_value_strings:
        actual_calibration_value_string = ""
        for i in range(len(calibration_value_string)):
            if calibration_value_string[i] in DIGITS:
                actual_calibration_value_string += calibration_value_string[i]
                break
            found = False
            for key, value in replacements.items():
                if calibration_value_string[i:i+len(key)] == key:
                    actual_calibration_value_string += replacements[key]
                    found = True
                    break
            if found:
                break

        for i in range(len(calibration_value_string)):
            j = len(calibration_value_string) - i - 1
            if calibration_value_string[j] in DIGITS:
                actual_calibration_value_string += calibration_value_string[j]
                break
            found = False
            for key, value in replacements.items():
                if calibration_value_string[j:j+len(key)] == key:
                    actual_calibration_value_string += replacements[key]
                    found = True
                    break
            if found:
                break
        
        calibration_values_sum += int(actual_calibration_value_string)
    print(calibration_values_sum)


def main():
    with open("resources/day01.txt", "r") as f:
        calibration_value_strings = f.read().rstrip().split("\n")
        part1(calibration_value_strings)
        part2(calibration_value_strings)
        


if __name__ == '__main__':
    main()