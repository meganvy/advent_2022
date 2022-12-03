import os
from collections import defaultdict

def part_one():
    script_dir = os.path.dirname(__file__)
    with open(os.path.join(script_dir, "input.txt"), "r") as f:
        val = 0
        for l in f:
            line = l.strip()
            first, second = list(line[:len(line)//2]), list(line[len(line)//2:])
            addend = 0
            for char in first:
                if char in second:
                    if char.upper() == char:
                        addend = (ord(char)-38)
                    else:
                        addend = (ord(char)-96)
                continue
            val += addend  
    return val


def part_two():
    script_dir = os.path.dirname(__file__)
    with open(os.path.join(script_dir, "input.txt"), "r") as f:
        val = 0
        badges = defaultdict(int)
        count = 0
        for l in f:
            # l o l
            line = list(set(l.strip()))
            for char in line:
                badges[char] += 1
            if count % 3 == 2:
                print(badges)
                addend = 0
                for k in badges.keys():
                    if badges[k] == 3:
                        if k.upper() == k:
                            addend = (ord(k)-38)
                        else:
                            addend = (ord(k)-96)
                print(addend)
                val += addend
                # reset dict
                badges = defaultdict(int)
            count += 1
    return val
                


# print(part_one())
print(part_two())