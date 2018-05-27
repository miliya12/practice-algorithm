package main

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{
			[]int{5, 2, 8, 6, 9, 7, 1, 4, 3},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, c := range cases {
		got := make([]int, len(c.in))
		copy(got, c.in)
		selectionSort(got)
		if !testEq(got, c.want) {
			t.Errorf("bubbleSort(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func testEq(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
