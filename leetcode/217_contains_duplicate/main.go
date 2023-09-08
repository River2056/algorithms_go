package containsduplicate

func ContainsDuplicate(nums []int) bool {
    set := make(map[int]int)
    for _, v := range nums {
        if _, ok := set[v]; ok {
            return true
        }
        set[v] = 1
    }
    return false
}
