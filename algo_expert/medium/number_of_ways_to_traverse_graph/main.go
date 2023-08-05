package numberofwaystotraversegraph

import "fmt"

func NumberOfWaysToTraverseGraph(width int, height int) int {
    cache := make(map[string]int, 0)
    return find(width, height, cache)
}

func find(width, height int, cache map[string]int) int {
    if width < 1 && height < 1 {
        return 0
    }

    if width == 1 || height == 1 {
        return 1
    }

    key1 := fmt.Sprintf("%d,%d", width, height)
    key2 := fmt.Sprintf("%d,%d", height, width)
    if _, ok := cache[key1]; ok {
        return cache[key1]
    }

    if _, ok := cache[key2]; ok {
        return cache[key2]
    }

    val := find(width - 1, height, cache) + find(width, height - 1, cache)
    cache[key1] = val

    return val
}
