from itertools import permutations, product
from functools import reduce
from tqdm import tqdm
import sys

def number_machine(numbers):
    new_numbers = []
    for number in numbers:
        if len(str(number)) % 2 == 0:
            str_number = str(number)
            index = int(len(str_number)/2)
            new_numbers.append(int(str(number)[:index]))
            new_numbers.append(int(str(number)[index:]))
        elif int(number) == 0:
            new_numbers.append(1)
        else:
            new_numbers.append(int(number)*2024)

    return new_numbers


with open(sys.argv[1], "r") as file1:
    # Reading from a file
    chars = []
    while line := file1.readline():
        numbers = line.rstrip().split(" ")
        for number in numbers:
            number = int(number)

    
for i in range(75):
    numbers = number_machine(numbers)
    print(i)
    print(len(numbers))
