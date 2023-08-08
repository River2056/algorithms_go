package kadanesalgorithm

import "testing"

func TestKadanesalgorithm(t *testing.T) {
	testCases := [][]interface{}{
		{[]int{3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4}, 19},
		{[]int{-1000, -1000, 2, 4, -5, -6, -7, -8, -2, -100}, 6},
	}

	for index, values := range testCases {
		param := values[0].([]int)
		expected := values[1].(int)

		result := KadanesAlgorithm(param)
		if result != expected {
			t.Errorf("%v. expected %v, got %v instead", index+1, expected, result)
		}
	}
}
