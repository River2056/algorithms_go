package romannumeralsdecoder

import "testing"

func TestRomanNumeralsDecoder(t *testing.T) {
	res := Decode("XXI")
	if res != 21 {
		t.Fatalf("expected 21, got %v\n", res)
	}

	res = Decode("I")
	if res != 1 {
		t.Fatalf("expected 1, got %v\n", res)
	}

	res = Decode("MMVIII")
	if res != 2008 {
		t.Fatalf("expected 2008, got %v\n", res)
	}

	res = Decode("MDCLXVI")
	if res != 1666 {
		t.Fatalf("expected 1666, got %v\n", res)
	}
}
