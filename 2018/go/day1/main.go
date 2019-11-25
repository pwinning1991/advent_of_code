package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	seen := map[int]bool{0: true}
	var nums []int

	s := bufio.NewScanner(file)

	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text, err)
		}
		nums = append(nums, n)
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	for {
		for _, n := range nums {
			sum += n
			if seen[sum] {
				fmt.Println(sum)
				return
			}
			seen[sum] = true
		}
	}

	fmt.Println(sum)

}
