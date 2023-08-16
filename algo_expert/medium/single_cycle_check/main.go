package singlecyclecheck

import "fmt"

func HasSingleCycleSelf(array []int) bool {
    cache := make(map[string]int)
    index := 0
    for {
        key := fmt.Sprintf("%v,%v", index, array[index])
        if _, ok := cache[key]; !ok {
            cache[key] = 1
        } else {
            if len(cache) != len(array) {
                return false
            }
            if index != 0 {
                return false
            }
            return true
        }
        jump := array[index]
        if jump < 0 {
            for i := 0; i != jump; i-- {
                index--
                if index < 0 {
                    index += len(array)
                }
            }
        } else if jump > 0 {
            for i := 0; i != jump; i++ {
                index++
                if index >= len(array) {
                    index %= len(array)
                }
            }
        }
    }
}

func HasSingleCycle(array []int) bool {
    count := 0
    index := 0
    for count < len(array) {
        if count > 0 && index == 0 {
            return false
        }
        count++
        index = getNextIdx(index, array)
    }
    return index == 0
}

func getNextIdx(index int, array []int) int {
    jump := array[index]
    nextIdx := (index + jump) % len(array)
    if nextIdx >= 0 {
        return nextIdx
    } else {
        return len(array) + nextIdx
    }
}
