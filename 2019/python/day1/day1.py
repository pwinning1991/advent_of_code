
def calculate_fuel(mass):
    fuel = mass//3 - 2
    return fuel

def read_input(filepath):
    inputs = [int(line.rstrip("\n")) for line in open(filepath)]
    return inputs

def main():
    total = 0
    nums = read_input('input.txt')
    for num in nums:
        total  += calculate_fuel(num)
    print(total)

    total = 0
    for num in nums:
        fuel = calculate_fuel(num)
        print(f"fuel calc after first call {fuel}")
        while fuel > 0:
            print(f"in the while loop, total is {total}")
            total += fuel
            print(f"added again {total}")
            fuel = calculate_fuel(fuel)
            print(f"current fuel calc {fuel}")
    print(f"total for part 2 {total}")


if __name__ == '__main__':
    main()

