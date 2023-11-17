package removeduplicatesfromsortedarray

func RemoveDuplicates(nums []int) int {
    count := 1
    j := 1

    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i - 1] {
            count++
        } else {
            count = 1
        }

        if count <= 2 {
            nums[j] = nums[i]
            j++
        }
    }

	return j
}
