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

	s := bufio.NewScanner(file)

	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("could not read %s: %v", s.Text, err)
		}
		sum += n
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)

}
