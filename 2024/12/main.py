from itertools import permutations, product
from functools import reduce
from tqdm import tqdm
import sys


with open(sys.argv[1], "r") as file1:
    # Reading from a file
    chars = []
    while line := file1.readline():
        numbers = line.rstrip().split(" ")
        for number in numbers:
            number = int(number)

