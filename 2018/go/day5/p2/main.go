package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := string(b)

	minLen := -1

	for u := range units(s) {
		r := react(trim(s, u))
		if minLen < 0 || len(r) < minLen {
			minLen = len(r)
		}
	}

	fmt.Println(minLen)

}

func trim(s string, b byte) string {
	const diff = 'a' - 'A'
	var r []byte
	for i := 0; i < len(s); i++ {
		if s[i] == b || s[i]+diff == b {
			continue
		}
		r = append(r, s[i])
	}

	return string(r)

}

func units(s string) map[byte]bool {
	s = strings.ToLower(s)
	u := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		u[s[i]] = true
	}
	return u
}

func react(s string) string {
	ok := true
	for ok {
		s, ok = step(s)
	}
	return s
}

func step(s string) (string, bool) {
	for i := 0; i < len(s)-1; i++ {
		if opposite(s[i], s[i+1]) {
			return s[:i] + s[i+2:], true
		}
	}
	return s, false
}

func opposite(a, b byte) bool {
	const diff = 'a' - 'A'
	return a+diff == b || b+diff == a
}
