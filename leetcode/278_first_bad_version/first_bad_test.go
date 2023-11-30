package firstbadversion

import (
	"fmt"
	"testing"
)

func TestFirstBadVersion(t *testing.T) {
	versions := []int{0, 0, 0, 1, 1}
	fmt.Println(Firstbadversion(len(versions)))
}
