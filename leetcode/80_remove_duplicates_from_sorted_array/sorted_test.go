package removeduplicatesfromsortedarray

import (
	"fmt"
	"testing"
)

func TestRemoveduplicatesfromsortedarray(t *testing.T) {
	fmt.Println(RemoveDuplicates([]int{1, 1, 1, 2, 2, 3}))          // 5 [1, 1, 2, 2, 3, _]
	fmt.Println(RemoveDuplicates([]int{0, 0, 1, 1, 1, 1, 2, 3, 3})) // 7 [0, 0, 1, 1, 2, 3, 3, _, _]
	fmt.Println(RemoveDuplicates([]int{1, 1, 1}))                   // 2 [1, 1, _]
}
