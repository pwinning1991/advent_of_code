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

	s := bufio.NewScanner(file)

	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %d, because of %v", n, err)
		}
		nums = append(nums, n)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	sumup := sum(nums)
	fmt.Printf("Here is the total sum %v", sumup)

}
