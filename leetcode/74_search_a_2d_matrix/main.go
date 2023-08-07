package searcha2dmatrix

func SearchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	length := m * n
	left, right := 0, length-1

	for left <= right {
		var mid int = (left + right) / 2
		var row int = mid / n
		var col int = mid % n
		element := matrix[row][col]
		if element == target {
			return true
		} else if element < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
