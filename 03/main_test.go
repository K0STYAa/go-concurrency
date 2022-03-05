package main

import "testing"

func TestSumSqr(t *testing.T) {
	data := []int{2, 4, 6, 8, 10}
	expected := 220
	actual := sumSquares(data)

	if expected != actual {
		t.Errorf("expected %d do not mutch actual %d", expected, actual)
	}
}
