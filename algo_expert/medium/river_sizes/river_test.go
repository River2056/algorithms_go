package riversizes

import (
	"fmt"
	"testing"
)

func TestRiver(t *testing.T) {
	matrix := [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
	}

	result := RiverSizes(matrix)
	fmt.Println(result) // [1, 2, 2, 2, 5]
}
