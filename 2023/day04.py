import re

import reader

def multispace_split(value):
    return re.split(r"\s+", value)

def get_cards(raw_cards):
    cards = []
    for raw_card in raw_cards:
        colon_index = raw_card.find(":")
        card_id = int(multispace_split(raw_card[:colon_index])[1])
        winning_numbers_raw, player_numbers_raw = raw_card[colon_index+2:].strip().split("|")
        winning_numbers = [int(number) for number in multispace_split(winning_numbers_raw.strip())]
        player_numbers = [int(number) for number in multispace_split(player_numbers_raw.strip())]
        cards.append({
            "id": card_id,
            "winning_numbers": winning_numbers,
            "player_numbers": player_numbers
        })
    return cards


def part1(cards):
    total_points = 0
    for card in cards:
        winning_numbers = set(card["winning_numbers"])\
            .intersection(set(card["player_numbers"]))
        num_winning_numbers = len(winning_numbers)
        if num_winning_numbers == 0:
            continue
        points = 2**(num_winning_numbers - 1)
        total_points += points

    # 20107
    print(total_points)



def part2(cards):
    for card in cards:
        winning_numbers = set(card["winning_numbers"])\
            .intersection(set(card["player_numbers"]))
        card["num_winning_numbers"] = len(winning_numbers)

    cards_to_process_ids = [card["id"] for card in cards]
    num_won_cards = len(cards)
    while cards_to_process_ids:
        card_to_process_id = cards_to_process_ids.pop()
        num_winning_numbers = cards[card_to_process_id-1]["num_winning_numbers"]
        for win_id in range(num_winning_numbers):
            cards_to_process_ids.append(card_to_process_id + win_id + 1)
        num_won_cards += num_winning_numbers

    # 8172507
    print(num_won_cards)




def main():
    raw_cards = reader.read_lines("resources/day04.txt")
    cards = get_cards(raw_cards)
    part1(cards)
    part2(cards)        


if __name__ == '__main__':
    main()