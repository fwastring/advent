from itertools import permutations, product
from functools import reduce
from tqdm import tqdm
import sys

with open(sys.argv[1], "r") as file1:
    # Reading from a file
    char_map = []
    results = []
    while line := file1.readline():
        row = line.rstrip().split()
        char_row = []
        for char in row:
            char_row.append(char)
        char_map.append(char_row)
    for row in char_map:
        print(row)

