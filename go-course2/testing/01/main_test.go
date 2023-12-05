package main

import "testing"

func TestAddSum(t *testing.T) {

	x := addNum(4, 9)
	if x != 13 {
		// t.Error("Expected 13, got", x)
		t.Error("Expected", 15, "Got", x)
	}
}
