package longestpalindromicstring

func longestPalindrome(s string) string {
	res := ""

	var find func(s, r string, left, right int)
	find = func(s, r string, left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if (right-left)+left+1 > len(r) {
				r = s[left : right+1]
			}
			left--
			right++
		}
	}

	for idx := range s {
		find(s, res, idx-1, idx+1)
		find(s, res, idx, idx+1)
	}

	return res
}
