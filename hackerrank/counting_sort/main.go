package countingsort

func CountingSort(arr []int) []int {
    res := make([]int, len(arr))
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	count := make([]int, max+1)
	for _, v := range arr {
		count[v]++
	}

    prefixSum := 0
    for i := 0; i <= max; i++ {
        
    }

	return []int{}
}
