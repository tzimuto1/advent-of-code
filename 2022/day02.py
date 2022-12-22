
RULES = {
    "ROCK": {
        "beats": "SCISSORS",
        "beaten_by": "PAPER",
        "points": 1 
    },
    "PAPER": {
        "beats": "ROCK",
        "beaten_by": "SCISSORS",
        "points": 2
    },
    "SCISSORS": {
        "beats": "PAPER",
        "beaten_by": "ROCK",
        "points": 3
    },
    "OUTCOMES": {
        "win": 6,
        "draw": 3,
        "lose": 0
    },
    "ALIASES": {
        "A": "ROCK",
        "B": "PAPER",
        "C": "SCISSORS"
    }

}


def get_score(my_play, opponent_play):
    score = RULES[my_play]["points"]
    if my_play == opponent_play :
        score += RULES["OUTCOMES"]["draw"]
    elif RULES[my_play]["beats"] == opponent_play:
        score += RULES["OUTCOMES"]["win"]
    else:
        score += RULES["OUTCOMES"]["lose"]
    return score

def part1(plays):
    # 12772
    outcome_encoding = {
        "X": "ROCK",
        "Y": "PAPER",
        "Z": "SCISSORS",
    }
    total_score = 0
    for play in plays:
        opponent_play = RULES["ALIASES"][play[0]]
        my_play = outcome_encoding[play[1]]
        total_score += get_score(my_play, opponent_play)

    print(total_score)

def part2(plays):
    # 11618
    outcome_encoding = {
        "X": "lose",
        "Y": "draw",
        "Z": "win"
    }
    total_score = 0
    for play in plays:
        opponent_play = RULES["ALIASES"][play[0]]
        my_play_encoding = outcome_encoding[play[1]]
        if my_play_encoding == "lose":
            my_play = RULES[opponent_play]["beats"]
        elif my_play_encoding == "draw":
            my_play = opponent_play
        else:
            my_play = RULES[opponent_play]["beaten_by"]
        total_score += get_score(my_play, opponent_play)
    print(total_score)



def main():
    with open("resources/day02.txt", "r") as f:
        plays = [play.split(" ") for play in f.read().strip().split("\n")]
        part1(plays)
        part2(plays)

if __name__ == '__main__':
    main()