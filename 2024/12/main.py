from itertools import permutations, product
from functools import reduce
from tqdm import tqdm
import sys

def explore_cluster(char_map, position, visited):
    i, j = position
    letter = char_map[i][j]
    cluster_size = 1
    fences = 4
    visited.add(position)

    for direction in [(-1, 0), (0, 1), (1, 0), (0, -1)]:
        next_pos = (i + direction[0], j + direction[1])
        if (0 <= next_pos[0] < len(char_map)) and (0 <= next_pos[1] < len(char_map[0])):
            next_letter = char_map[next_pos[0]][next_pos[1]]
            if next_letter == letter:
                fences -= 1
                if next_pos not in visited:
                    sub_cluster_size, sub_fences = explore_cluster(char_map, next_pos, visited)
                    cluster_size += sub_cluster_size
                    fences += sub_fences

    return cluster_size, fences


with open(sys.argv[1], "r") as file1:
    char_map = []
    visited = set()

    while line := file1.readline():
        char_map.append(list(line.rstrip()))

    sum = 0
    for i in range(len(char_map)):
        for j in range(len(char_map[i])):
            position = (i, j)

            if position not in visited:
                cluster_size, fences = explore_cluster(char_map, position, visited)
                letter = char_map[i][j]
                sum += cluster_size*fences
                print(letter, cluster_size, fences)
    print(sum)

