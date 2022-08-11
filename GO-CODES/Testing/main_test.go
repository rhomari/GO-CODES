package main

import "testing"

func TestArraySum(t *testing.T) {
	var sum int64 = 0
	sum = ArraySum(1, 2, 3, 4, 5)
	if sum != 15 {
		t.Errorf("Expected 15, got %d", sum)
	}
}
