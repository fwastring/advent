from itertools import permutations, product
from functools import reduce
from tqdm import tqdm
import sys

def check_score(char_map, position, number, visited):
    if number == 9:
        if position not in visited:
            # visited.add(position)
            return 1  # Found a valid path
    count = 0
    for direction in [(-1, 0), (0, 1), (1, 0), (0, -1)]:
        next_pos = (position[0] + direction[0], position[1] + direction[1])
        if (0 <= next_pos[0] < len(char_map)) and (0 <= next_pos[1] < len(char_map[0])):  # Ensure bounds
            if char_map[next_pos[0]][next_pos[1]] == number + 1:  # Continue the sequence
                visited.add(position)
                count += check_score(char_map, next_pos, number + 1, visited)
                visited.remove(position)
    return count

with open(sys.argv[1], "r") as file1:
    # Reading from a file
    char_map = []
    results = []
    while line := file1.readline():
        row = str(line.rstrip())
        chars = []
        for char in row:
           chars.append(int(char))
        char_map.append(chars)

    sum = 0
    for i in range(len(char_map)):
        for j in range(len(char_map[i])):
            if char_map[i][j] == 0:
                sum += (check_score(char_map, (i,j), 0, {(i,j)}))
    print(sum)


