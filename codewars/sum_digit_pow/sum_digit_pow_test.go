package sumdigitpow

import (
	"fmt"
	"testing"
)

func TestSumdigitpow(t *testing.T) {
    res := SumDigitPow(1, 10)
    fmt.Println(res) // [1, 2, 3, 4, 5, 6, 7, 8, 9]

    res = SumDigitPow(1, 100)
    fmt.Println(res) // [1, 2, 3, 4, 5, 6, 7, 8, 9, 89]

    res = SumDigitPow(12157692622039623538, 12157692622039623539)
    fmt.Println(res) // [12157692622039623539]
}
