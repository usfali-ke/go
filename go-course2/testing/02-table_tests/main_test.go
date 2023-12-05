package main

import "testing"

type test struct {
	data   []int
	answer int
}

func TestAddSum(t *testing.T) {

	tests := []test{
		test{[]int{12, 1}, 13},
		test{[]int{1, 1}, 2},
		test{[]int{6, 70}, 76},
		test{[]int{12, 28}, 40},
	}

	for _, v := range tests {
		x := addNum(v.data...)
		if x != v.answer {
			// t.Error("Expected 13, got", x)
			t.Error("Expected", v.answer, "Got", x)
		}

	}
}
