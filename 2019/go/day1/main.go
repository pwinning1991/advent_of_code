package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func calculateFuel(mass int) int {
	fuel := (mass / 3) - 2

	return fuel

}

func calculateTotalFuel(mass int) int {
	var total int
	newFuel := calculateFuel(mass)
	for newFuel >= 0 {
		total += newFuel
		newFuel = calculateFuel(newFuel)
	}
	return total
}

func sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Not able to read file %v", file)
	}
	defer file.Close()

	var nums []int
	var nums2 []int

	s := bufio.NewScanner(file)

	for s.Scan() {
		var n int
		var n2 int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %d, because of %v", n, err)
		}
		n2 = calculateTotalFuel(n)
		n = calculateFuel(n)
		nums = append(nums, n)
		nums2 = append(nums2, n2)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	sumup := sum(nums)
	fmt.Printf("Here is the total sum %v\n", sumup)
	sump2 := sum(nums2)
	fmt.Printf("Here is the total for part 2: %v\n", sump2)

}
