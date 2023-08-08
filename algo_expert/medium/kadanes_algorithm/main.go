package kadanesalgorithm

import (
	"fmt"
	"math"
)

func KadanesAlgorithmSelf(array []int) int {
	if len(array) == 0 {
		return -1
	}
	if len(array) == 1 {
		return array[0]
	}

	max := 0
	current_count := array[0]
	for i := 1; i < len(array); i++ {
		fmt.Printf("current_count: %v, max: %v\n", current_count, max)
		if current_count < 0 {
			current_count = 0
		}
		current_count += array[i]
		if current_count >= max {
			max = current_count
		}

	}

	if max <= 0 {
		max = -1
	}

	return max
}

func KadanesAlgorithm(array []int) int {
	maxEndingHere := float64(array[0])
	maxSoFar := array[0]
	for i := 1; i < len(array); i++ {
        maxEndingHere = math.Max(float64(array[i]), float64(array[i] + int(maxEndingHere)))
		if int(maxEndingHere) > maxSoFar {
			maxSoFar = int(maxEndingHere)
		}
	}

	if maxSoFar <= 0 {
		maxSoFar = -1
	}

	return maxSoFar
}
