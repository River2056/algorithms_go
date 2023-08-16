package combinationsum

import (
	"fmt"
	"testing"
)

func TestCombinationsum(t *testing.T) {
	result := combinationSum([]int{2, 3, 4}, 6)
	fmt.Printf("result: %v\n", result) // {{2, 2, 3}, {7}}
}
