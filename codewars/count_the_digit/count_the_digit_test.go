package countthedigit

import (
	"fmt"
	"testing"
)

func TestCountthedigit(t *testing.T) {
	res := NbDig(10, 1)
	fmt.Println(res) // 4
	if res != 4 {
		t.Errorf("expected 4, got %v instead\n", res)
	}

	res = NbDig(5750, 0)
	fmt.Println(res) // 4700
	if res != 4700 {
		t.Errorf("expected 4700, got %v instead\n", res)
	}
}
