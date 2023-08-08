package searchinrotatedsortedarray

func Search(nums []int, target int) int {
	var pivot int
	left, right := 0, len(nums)-1
	for left < right {
		var mid int = left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] <= nums[right] {
			right = mid
		}
	}
	pivot = left

	// search left
	if target >= nums[pivot] && target <= nums[len(nums)-1] {
		return binarySearch(nums, pivot, len(nums)-1, target)
	} else {
		// search right
		return binarySearch(nums, 0, pivot, target)
	}
}

func binarySearch(nums []int, left, right, target int) int {
	for left <= right {
		var mid int = left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
