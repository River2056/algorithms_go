package maximumsubarray

import "math"

func MaxSubArray(nums []int) int {
	current := nums[0]
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if current < nums[i] {
			current = nums[i]
		} else {
			current += nums[i]
		}
		max = int(math.Max(float64(max), float64(current)))
	}

	return max
}
