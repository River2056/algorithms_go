package arraydiff

func ArrayDiff(a, b []int) []int {
    m := make(map[int]bool)
    for _, v := range b {
        m[v] = true
    }

    res := make([]int, 0)
    for _, v := range a {
        if _, ok := m[v]; !ok {
            res = append(res, v)
        }
    }

    return res
}
