package main

import (
	"fmt"
	"testing"
)

func TestIsSafe(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		want bool
	}{
		{0, 0, true},
		{99, 99, false},
		{51, 7, true},
		{59, 75, false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%v %v", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) {
			ans := isSafe(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestSumDigits(t *testing.T) {
	var tests = []struct {
		a    int
		want int
	}{
		{0, 0},
		{10, 1},
		{11, 2},
		{66, 12},
		{111, 3},
		{9999, 36},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%v", tt.a)
		t.Run(testName, func(t *testing.T) {
			ans := sumDigits(tt.a)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func BenchmarkMain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		main()
	}
}
