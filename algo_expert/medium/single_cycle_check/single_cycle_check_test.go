package singlecyclecheck

import "testing"

func TestSinglecyclecheck(t *testing.T) {
	result := HasSingleCycle([]int{2, 3, 1, -4, -4, 2})
	if !result {
		t.Errorf("expected true, got %v instead\n", result)
	}
}
