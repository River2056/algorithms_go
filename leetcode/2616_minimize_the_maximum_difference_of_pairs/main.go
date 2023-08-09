package minimizethemaximumdifferenceofpairs

import (
	"math"
	"sort"
)

func MinimizeMax(nums []int, p int) int {
    if p == 0 {
        return 0
    }

    sort.Slice(nums, func(a, b int) bool {
        return nums[a] < nums[b]
    })
    
    var res int
    left, right := 0, nums[len(nums) - 1] - nums[0]
    res = right
    for left <= right {
        var mid int = left + (right - left) / 2
        if isValid(nums, p, mid) {
            res = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return res
}

func isValid(nums []int, p, threshold int) bool {
    i, cnt := 0, 0
    for i < len(nums) - 1 {
        if math.Abs(float64(nums[i]) - float64(nums[i + 1])) <= float64(threshold) {
            cnt++
            i += 2
        } else {
            i++
        }
        if cnt == p {
            return true
        }
    }

    return false
}
