# apologies in advance i'm waiting until a weekend to clean up / make helpers for reading inputs

import os

def part_one():
    max_calories = 0
    script_dir = os.path.dirname(__file__)
    with open(os.path.join(script_dir, "input.txt"), "r") as f:
        val = 0
        for line in f:
            if line != "\n":
                val += int(line)
            else:
                max_calories = max(max_calories, val)
                val = 0
    return max_calories


def part_two():
    elves = {}
    script_dir = os.path.dirname(__file__)
    with open(os.path.join(script_dir, "input.txt"), "r") as f:
        carrying = 0
        for line in f:
            if line != "\n":
                carrying += int(line)
            else:
                elves[carrying] = 0
                carrying = 0
    elves[carrying] = 0
    count = 0
    top_three = sorted(elves.keys())[-3:]
    return sum(top_three)


print(part_one())
print(part_two())