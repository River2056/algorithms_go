package levenshteindistance

import (
	"strings"
)

func LevenshteinDistance(a, b string) int {
    if len(strings.TrimSpace(a)) == 0 {
        return len(b)
    }

    if len(strings.TrimSpace(b)) == 0 {
        return len(a);
    }

	a = " " + a
	b = " " + b
	dp := make([][]int, len(a))
    for i := range a {
		dp[i] = make([]int, len(b))
	}

	for i := 0; i < len(b); i++ {
		dp[0][i] = i
	}

	for i := 0; i < len(a); i++ {
		dp[i][0] = i
	}

	for i := 1; i < len(a); i++ {
		for j := 1; j < len(b); j++ {
			if a[i] == b[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// neighbours
				values := []int{dp[i-1][j-1], dp[i-1][j], dp[i][j-1]}
				min := values[0]
				for _, v := range values {
					if v < min {
						min = v
					}
				}
				dp[i][j] = 1 + min
			}
		}
	}

	return dp[len(dp)-1][len(dp[len(dp)-1])-1]
}
