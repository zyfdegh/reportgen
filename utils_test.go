package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestTop(t *testing.T) {
	var cases = []struct {
		a []int
		n int
		b []int
	}{
		{[]int{2, 4, 6, 3, 1}, 2, []int{4, 6}},
		{[]int{2, 4, 6, 4, 2}, 3, []int{2, 4, 6}},
		{[]int{18, 30, 23, 9, 8, 8, 2, 8, 4, 1, 12, 3, 4, 2, 3, 4, 4, 1, 4, 26, 1, 6}, 5, []int{12, 18, 23, 26, 30}},
	}

	for _, c := range cases {
		got := top(c.a, c.n)
		assert.Equal(t, c.b, got)
	}
}

func TestHour(t *testing.T) {
	var cases = []struct {
		f float32
		h int
	}{
		{0.004166666666, 0},
		{0.375, 9},
		{0.389583333333, 9},
		{0.404861111111, 9},
		{0.415277777777, 9},
		{0.417361111111, 10},
		{0.979166666666, 23},
	}

	for _, c := range cases {
		got := hour(c.f)
		assert.Equal(t, c.h, got)
	}
}