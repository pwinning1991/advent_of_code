#!/usr/bin/env python3
import re
from collections import defaultdict

data = [x.strip() for x in open('input.txt')]

claims = map(lambda s: map(int, re.findall(r'-?\d+', s)), data)
m = defaultdict(list)
overlaps = {}
