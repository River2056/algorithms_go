package numberof1bits

import (
	"strconv"
)

func HammingWeight(num uint32) int {
    s := strconv.FormatInt(int64(num), 2)
    count := 0
    for _, c := range s {
        if c == '1' {
            count++
        }
    }
    return count
}
