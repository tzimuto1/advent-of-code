
from collections import namedtuple
Game = namedtuple("Game", "id picks")

def get_games(raw_games):
    for raw_game in raw_games:
        colon_index = raw_game.find(":")



def get_games(raw_games):
    games = []
    for raw_game in raw_games:
        colon_index = raw_game.find(":")
        game_id = int(raw_game[:colon_index].split(" ")[1])
        game_picks = raw_game[colon_index+2:].replace(";", ",").split(", ")
        games.append({
            "id": game_id,
            "picks": game_picks
        })
    return games

def part1(raw_games):
    # 2720
    color_counts = {
        "red": 12,
        "green": 13,
        "blue": 14
    }

    sum_of_possible_games = 0

    games = get_games(raw_games)

    for game in games:
        game_id = game["id"]
        game_picks = game["picks"]
        is_game_possible = True
        for pick in game_picks:
            count, color = pick.split(" ")
            count = int(count)
            if count > color_counts[color]:
                is_game_possible = False
                break
        if is_game_possible:
            sum_of_possible_games += game_id
    print(sum_of_possible_games)


def part2(raw_games):
    games = get_games(raw_games)

    sum_of_cubes = 0

    for game in games:
        color_maximums = {
            "red": 0,
            "green": 0,
            "blue": 0
        }
        for pick in game["picks"]:
            count, color = pick.split(" ")
            count = int(count)
            if color_maximums[color] < count:
                color_maximums[color] = count
        cube = color_maximums["red"] * color_maximums["green"] * color_maximums["blue"]
        sum_of_cubes += cube

    print(sum_of_cubes)




def main():
    with open("resources/day02.txt", "r") as f:
        raw_games = f.read().rstrip().split("\n")
        part1(raw_games)
        part2(raw_games)
        


if __name__ == '__main__':
    main()
