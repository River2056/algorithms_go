package countingsort

import (
	"fmt"
	"testing"
)

func TestCountingSort(t *testing.T) {
	arr := []int{3, 5, 0, 2, 1, 4, 8, 1, 0, 0, 2}

	sorted := CountingSort(arr)
	fmt.Println(sorted)
}
