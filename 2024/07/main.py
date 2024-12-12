from itertools import permutations, product
from functools import reduce
from tqdm import tqdm

def eval_left_to_right(expression, max_value):
    tokens = expression.split()  # Split the expression into tokens
    result = int(tokens[0])    # Start with the first number
    
    # Process the tokens left to right
    i = 1
    while i < len(tokens):
        operator = tokens[i]
        operand = int(tokens[i + 1])

        
        if operator == '+':
            result += operand
        elif operator == '*':
            result *= operand
        elif operator == '||':
            result = int(str(result) + str(operand))
        else:
            raise ValueError(f"Unsupported operator: {operator}")
        
        if result > max_value:
            break
        
        i += 2  # Move to the next operator-operand pair
    
    return result

with open("input", "r") as file1:
    # Reading from a file
    results = []
    while line := file1.readline():
        print(line)
        parsed = line.rstrip()
        split = parsed.split(":")
        result = int(split[0].rstrip())
        arguments = list(map(int, split[1].rstrip().split(" ")[1:]))
        ops = ['+', '*', "||"]
        for all_ops in list(product(ops, repeat=len(arguments)-1)):
            expr = f"{arguments[0]}"
            for i in range(1, len(arguments)):
                expr += f" {all_ops[i-1]} {arguments[i]}"
            value = eval_left_to_right(expr, result)
            if value == result:
                results.append(value)
                break
        

    print(results)
    print(reduce(lambda a, b: a + b, results))

