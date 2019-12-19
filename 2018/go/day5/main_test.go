package main

import (
	"fmt"
	"testing"
)

func TestSteps(t *testing.T) {
	tt := []struct {
		in, out string
		ok      bool
	}{
		{"aA", "", true},
		{"aa", "aa", false},
		{"aAbB", "bB", true},
		{"aAa", "a", true},
	}

	for _, tc := range tt {
		t.Run(tc.in, func(t *testing.T) {
			out, ok := step(tc.in)
			if tc.ok != ok {
				t.Errorf("expected ok to be %v; got %v", ok, tc.ok)
			}
			if tc.out != out {
				t.Errorf("expected ok to be %v; got %v", out, tc.out)
			}

		})
	}

}

func TestOpposite(t *testing.T) {
	tt := []struct {
		a, b byte
		res  bool
	}{
		{'a', 'A', true},
		{'A', 'a', true},
		{'a', 'a', false},
		{'A', 'A', false},
		{'b', 'A', false},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%c-%c", tc.a, tc.b), func(t *testing.T) {
			if res := opposite(tc.a, tc.b); res != tc.res {
				t.Fatalf("expected opposite of %c and %c to be %v, got %v", tc.a, tc.b, tc.res, res)
			}

		})
	}
}
