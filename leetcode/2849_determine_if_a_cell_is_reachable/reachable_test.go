package determineifacellisreachable

import (
	"fmt"
	"testing"
)

func TestDetermineifacellisreachable(t *testing.T) {
	fmt.Println(IsReadhableAtTime(2, 4, 7, 7, 6)) // true
	fmt.Println(IsReadhableAtTime(3, 1, 7, 3, 3)) // false
}
