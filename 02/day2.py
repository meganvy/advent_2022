import os

# i DON'T wanna talk about it actually
def part_one():
    script_dir = os.path.dirname(__file__)
    with open(os.path.join(script_dir, "input.txt"), "r") as f:
        total_score = 0
        for line in f:
            score = 0
            plays = line.strip("\n").split(" ")
            if plays[1] == "X":
                score += 1
                if plays[0] == "A":
                    score += 3
                elif plays[0] == "B":
                    score += 0
                else:
                    score += 6
            elif plays[1] == "Y":
                score += 2
                if plays[0] == "A":
                    score += 6
                elif plays[0] == "B":
                    score += 3
                else:
                    score += 0
            else:
                score += 3
                if plays[0] == "A":
                    score += 0
                elif plays[0] == "B":
                    score += 6
                else:
                    score += 3
            total_score += score
    return total_score


def part_two():
    script_dir = os.path.dirname(__file__)
    wins = {"A": "Z", "B": "X", "C": "Y"}
    losses = {"A": "Y", "B": "Z", "C": "X"}
    draws = {"A": "X", "B": "Y", "C": "Z"}
    scores = {"X": 1, "Y": 2, "Z": 3}
    with open(os.path.join(script_dir, "input.txt"), "r") as f:
        total_score = 0
        for line in f:
            score = 0
            plays = line.strip("\n").split(" ")
            if plays[1] == "X":
                score += 0
                score += scores[wins[plays[0]]]
            elif plays[1] == "Y":
                score += 3
                score += scores[draws[plays[0]]]
            else: 
                score += 6
                score += scores[losses[plays[0]]]
            total_score += score
    return total_score



print(part_one())
print(part_two())