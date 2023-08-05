package numberofwaystotraversegraph

import "testing"

func TestNumberofwaystotraversegraph(t *testing.T) {
	testCases := [][]interface{}{
		[]interface{}{[]int{4, 3}, 10},
	}

	for _, values := range testCases {
		param := values[0].([]int)
		expected := values[1].(int)
		result := NumberOfWaysToTraverseGraph(param[0], param[1])
		if result != expected {
			t.Errorf("expected %v, got %v instead", expected, result)
		}
	}
}
