package stableinternships

import (
	"sort"
	"testing"
)

func TestStableinternships(t *testing.T) {
	testCases := [][]interface{}{
		{[][]int{{0, 1, 2}, {1, 0, 2}, {1, 2, 0}}, [][]int{{2, 1, 0}, {1, 2, 0}, {0, 2, 1}}, [][]int{{0, 0}, {1, 1}, {2, 2}}},
		{[][]int{{0, 1, 2, 3}, {2, 1, 3, 0}, {0, 2, 3, 1}, {3, 1, 0, 2}}, [][]int{{1, 3, 2, 0}, {0, 1, 2, 3}, {1, 2, 3, 0}, {3, 0, 2, 1}}, [][]int{{0, 1}, {1, 2}, {2, 0}, {3, 3}}},
	}

	for idx, values := range testCases {
		interns := values[0].([][]int)
		teams := values[1].([][]int)
		expected := values[2].([][]int)

		result := StableInternships(interns, teams)
		if !equals(result, expected) {
			t.Errorf("%v. expected %v, got %v instead\n", idx+1, expected, result)
		}
	}
}

func equals(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if len(v) != len(b[i]) {
			return false
		}
	}

	aArray := make([]int, 0)
	for _, v := range a {
		for _, element := range v {
			aArray = append(aArray, element)
		}
	}

	bArray := make([]int, 0)
	for _, v := range b {
		for _, element := range v {
			bArray = append(bArray, element)
		}
	}

	sort.Slice(aArray, func(a1, a2 int) bool {
		return aArray[a1] < aArray[a2]
	})

	sort.Slice(bArray, func(b1, b2 int) bool {
		return bArray[b1] < bArray[b2]
	})

	for i := range aArray {
		if aArray[i] != bArray[i] {
			return false
		}
	}

	return true
}
