package longestpalindromicstring

import (
	"fmt"
	"testing"
)

func TestLongestpalindromicstring(t *testing.T) {
    result := longestPalindrome("babad")
    fmt.Println(result)

    result = longestPalindrome("cbbd")
    fmt.Println(result)
}
