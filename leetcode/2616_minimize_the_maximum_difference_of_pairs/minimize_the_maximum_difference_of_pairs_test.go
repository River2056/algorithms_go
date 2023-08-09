package minimizethemaximumdifferenceofpairs

import "testing"

func TestMinimizethemaximumdifferenceofpairs(t *testing.T) {
	testCases := [][]interface{}{
		{[]int{10, 1, 2, 7, 1, 3}, 2, 1},
        {[]int{4, 2, 1, 2}, 1, 0},
	}

	for idx, values := range testCases {
		nums := values[0].([]int)
		p := values[1].(int)
		expected := values[2].(int)

		result := MinimizeMax(nums, p)
		if result != expected {
			t.Errorf("%v. expected %v, got %v instead\n", idx+1, expected, result)
		}
	}
}
