package riversizes

func RiverSizes(matrix [][]int) []int {
    res := make([]int, 0)
    visited := make([][]bool, len(matrix))
    for idx, innerArray := range matrix {
        visited[idx] = make([]bool, len(innerArray))
    }

    var dfs func(x, y int, count *int)
    dfs = func(x, y int, count *int) {
        if y >= len(matrix) || y < 0 {
            return
        }
        arr := matrix[y]
        if x >= len(arr) || x < 0 {
            return
        }

        if matrix[y][x] == 0 || visited[y][x] {
            return
        }

        if matrix[y][x] == 1 {
            *count++
            visited[y][x] = true
        }

        dfs(x + 1, y, count)
        dfs(x - 1, y, count)
        dfs(x, y + 1, count)
        dfs(x, y - 1, count)
    }

    for y := range matrix {
        for x := range matrix[y] {
            count := 0
            dfs(x, y, &count)
            if count > 0 {
                res = append(res, count)
            }
        }
    }

	return res
}
