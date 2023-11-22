package twosumarrayissorted

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	fmt.Println(TwoSum([]int{2, 7, 11, 15}, 9)) // [1, 2]
	fmt.Println(TwoSum([]int{2, 3, 4}, 6))      // [1, 2]
	fmt.Println(TwoSum([]int{0, -1}, -1))       // [1, 2]
}
