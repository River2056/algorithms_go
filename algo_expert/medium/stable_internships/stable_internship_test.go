package stableinternships

import "testing"

func TestStableinternships(t *testing.T) {
    testCases := [][]interface{}{
        {[][]int {{0, 1, 2}, {1, 0, 2}, {1, 2, 0}}, [][]int {{2, 1, 0}, {1, 2, 0}, {0, 2, 1}}, [][]int {{0, 0}, {1, 1}, {2, 2}}},
    }

    for idx, values := range testCases {
        interns := values[0].([][]int)
        teams := values[1].([][]int)
        expected := values[2].([][]int)

        result := StableInternships(interns, teams)
        if !equals(result, expected) {
            t.Errorf("%v. expected %v, got %v instead\n", idx + 1, expected, result)
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
        
        for j, element := range a[i] {
            if element != b[i][j] {
                return false
            }
        }
    }

    return true
}
