#!/usr/bin/env python3
from collections import Counter
nums_to_multiply = {'number_of_2s': 0, 'number_of_3s': 0}

data = [x.strip() for x in open('input.txt')]

for line in data:
    letter_counter = Counter()
    for s in line:
        letter_counter[s] += 1
    has_two = False
    has_three = False
    for k,v in letter_counter.items():
        if v == 2:
            has_two = True
        elif v == 3:
            has_three = True
            nums_to_multiply['number_of_3s'] += 1
    if has_two:
        nums_to_multiply['number_of_2s'] += 1
    elif has_two:
        nums_to_multiply['number_of_3s'] += 1


print(nums_to_multiply['number_of_2s'] * nums_to_multiply['number_of_3s'])

