package validpalindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    sTemp := make([]rune, 0)
    for _, c := range s {
        if unicode.IsLetter(c) || unicode.IsDigit(c) {
            sTemp = append(sTemp, c)
        }
    }

    left, right := 0, len(sTemp) - 1
    for left <= right {
        l := sTemp[left]
        r := sTemp[right]
        if l != r {
            return false
        } else {
            left++
            right--
        }
    }
    return true
}
