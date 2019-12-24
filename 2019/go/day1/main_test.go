package main

import (
	"fmt"
	"testing"
)

func TestCalculateFuel(t *testing.T) {

	cases := []struct {
		mass   int
		answer int
	}{
		{mass: 12, answer: 2},
		{mass: 14, answer: 2},
		{mass: 1969, answer: 654},
		{mass: 100756, answer: 33583},
	}

	for _, v := range cases {
		t.Run(fmt.Sprintf("testing with the mass of %d", v.mass), func(t *testing.T) {
			got := calculateFuel(v.mass)
			want := v.answer
			assertEqual(t, got, want)

		})
	}

}

func TestSum(t *testing.T) {
	cases := []struct {
		nums []int
		sum  int
	}{
		{nums: []int{1, 2, 3, 4, 5}, sum: 15},
		{nums: []int{5, 2, 3, 4, 5}, sum: 19},
		{nums: []int{2, 4, 6, 8}, sum: 20},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("Testing sum function on %v", tc.nums), func(t *testing.T) {
			got := sum(tc.nums)
			want := tc.sum
			assertEqual(t, got, want)

		})
	}

}

func TestCalculateTotalFuel(t *testing.T) {

	tc := []struct {
		mass  int
		total int
	}{
		{mass: 14, total: 2},
		{mass: 1969, total: 966},
		{mass: 100756, total: 50346},
	}

	for _, tt := range tc {
		t.Run(fmt.Sprintf("Running test total fuel on %v", tt.mass), func(t *testing.T) {
			got := calculateTotalFuel(tt.mass)
			want := tt.total
			assertEqual(t, got, want)

		})
	}

}

func assertEqual(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
