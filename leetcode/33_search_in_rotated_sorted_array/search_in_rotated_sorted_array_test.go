package searchinrotatedsortedarray

import "testing"

func TestSearchinrotatedsortedarray(t *testing.T) {
	// 5, 6, 7, 0, 1, 2, 4
	testCases := [][]interface{}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{1}, 1, 0},
		{[]int{1, 3}, 2, -1},
		{[]int{5, 1, 3}, 1, 1},
	}

	for idx, values := range testCases {
		inputArray := values[0].([]int)
		target := values[1].(int)
		expected := values[2].(int)

		result := Search(inputArray, target)
		if result != expected {
			t.Errorf("%v. expected %v, got %v instead", idx+1, expected, result)
		}
	}
}
