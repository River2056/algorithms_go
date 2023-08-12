package combinationsum

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	var dfs func(i, total int, cur *[]int)
	dfs = func(i, total int, cur *[]int) {
		if total == target {
			copyArray := make([]int, len(*cur))
			copy(copyArray, *cur)
			res = append(res, copyArray)
			return
		}

		if i >= len(candidates) || total > target {
			return
		}

		*cur = append(*cur, candidates[i])
		dfs(i, total+candidates[i], cur)
		*cur = (*cur)[:len(*cur)-1]
		dfs(i+1, total, cur)
	}

	cur := make([]int, 0)
	dfs(0, 0, &cur)

	return res
}
