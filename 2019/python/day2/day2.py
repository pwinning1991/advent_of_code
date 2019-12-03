import sys
def run_opscode(list_of_nums, pc):
    if list_of_nums[pc] == 99:
        return False
    opcode, input_one, input_two, output = list_of_nums[pc:pc+4]
    if opcode == 1:
        list_of_nums[output] = list_of_nums[input_one] + list_of_nums[input_two]
        return True
    elif opcode == 2:
        list_of_nums[output] = list_of_nums[input_one] * list_of_nums[input_two]
        return True

def read_input(filepath):
    return [int(x) for x in open(filepath).read().split(",")]

def run_program(list_of_nums, pc=0):
    while run_opscode(list_of_nums, pc):
        pc += 4


def main():
    list_of_nums = read_input('input.txt')
    list_of_nums1 = list_of_nums.copy()
    list_of_nums1[1], list_of_nums1[2] = 12,2
    print(list_of_nums1)
    run_program(list_of_nums1)
    print(f"part1: {list_of_nums1[0]}")

    for x in range(100):
        for y in range(100):
            list_of_nums1 = list_of_nums.copy()
            list_of_nums1[1], list_of_nums1[2] = x, y 
            run_program(list_of_nums1)

            if list_of_nums1[0] == 19690720:
                result = 100 * x + y
                print(f"Answer to part 2: {result}")
                sys.exit()



if __name__ == '__main__':
    main()
