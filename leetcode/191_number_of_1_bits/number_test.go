package numberof1bits

import (
	"fmt"
	"testing"
)

func TestNumberOfBits(t *testing.T) {
	fmt.Println(HammingWeight(11))         // 3
	fmt.Println(HammingWeight(128))        // 1
	fmt.Println(HammingWeight(4294967293)) // 31
}
