package validpalindrome

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T) {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome(" "))                              // true
	fmt.Println(isPalindrome("0P"))                             // false
}
