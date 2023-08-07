package searcha2dmatrix

import "testing"

func TestSearcha2dmatrix(t *testing.T) {
	testCases := [][]interface{}{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
		{[][]int{{1}}, 2, false},
		{[][]int{{1, 3}}, 3, true},
		{[][]int{{3}}, 4, false},
	}

	for index, values := range testCases {
		param := values[0].([][]int)
		target := values[1].(int)
		expected := values[2].(bool)

		result := SearchMatrix(param, target)
		if result != expected {
			t.Errorf("%v. expected %v, got %v instead", index+1, expected, result)
		}
	}
}
