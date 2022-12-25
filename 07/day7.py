# apologies in advance i'm waiting until a weekend to clean up / make helpers for reading inputs

import os

class node:
    def __init__(self, name):
        self.parent = None
        self.name = name
        self.subdirs = []


def part_one():
    n = node(name="/")
    root = node
    script_dir = os.path.dirname(__file__)
    with open(os.path.join(script_dir, "input.txt"), "r") as f:
        val = 0
        for line in f:
            print(line)
            print(n)
            if not line.startswith("$ "): # files and directories
                newVal = n.subdirs
                n.subdirs = newVal.append(node(line[5:], [], n))
            else:
                if line.startswith("$ cd") and line[5:] != "/":
                    for s in n.subdirs:
                        if s.name == line[5:]:
                            n = s
                    else:
                        n = n.parent
    # print(root)



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
# print(part_two())