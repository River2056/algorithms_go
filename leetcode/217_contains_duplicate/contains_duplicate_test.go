package containsduplicate

import "testing"

func TestContainsduplicate(t *testing.T) {
	res := ContainsDuplicate([]int{1, 2, 3, 1})
	if !res {
		t.Errorf("expected true, got %v instead\n", res)
	}

	res = ContainsDuplicate([]int{1, 2, 3, 4})
	if res {
		t.Errorf("expected false, got %v instead\n", res)
	}

	res = ContainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
	if !res {
		t.Errorf("expected true, got %v instead\n", res)
	}
}
