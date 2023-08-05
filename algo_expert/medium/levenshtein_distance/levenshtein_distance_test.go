package levenshteindistance

import "testing"

func TestLevenshteinDistance(t *testing.T) {
	result := LevenshteinDistance("abc", "yabd")
	if result != 2 {
		t.Errorf("expected 2, got %v instead", result)
	}

	result = LevenshteinDistance("", "abc")
	if result != 3 {
		t.Errorf("expected 3, got %v instead", result)
	}

	result = LevenshteinDistance("algoexpert", "algozexpert")
	if result != 1 {
		t.Errorf("expected 1, got %v instead", result)
	}

	result = LevenshteinDistance("abc", "abcx")
	if result != 1 {
		t.Errorf("expected 1, got %v instead", result)
	}

	result = LevenshteinDistance("cereal", "saturdzz")
	if result != 7 {
		t.Errorf("expected 7, got %v instead", result)
	}
}
