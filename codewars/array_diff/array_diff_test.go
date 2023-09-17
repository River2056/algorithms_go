package arraydiff

import (
	"fmt"
	"testing"
)

func TestArraydiff(t *testing.T) {
	res := ArrayDiff([]int{1, 2}, []int{1})
	fmt.Println(res) // [2]

	res = ArrayDiff([]int{1, 2, 2, 2, 3}, []int{2})
	fmt.Println(res) // [1, 3]
}
