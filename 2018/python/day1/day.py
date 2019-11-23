#!/usr/bin/env python3
import itertools

nums = [int(x) for x in open('input.txt')]

#part1
print(sum(nums))


#part2
counter = 0
seen = {0}

for num in itertools.cycle(nums):
    counter += num
    if counter in seen:
        print(counter)
        break
    seen.add(counter)
