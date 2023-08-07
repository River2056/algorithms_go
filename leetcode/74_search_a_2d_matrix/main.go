package searcha2dmatrix

import "fmt"

func SearchMatrix(matrix [][]int, target int) bool {
    rows := make([]int, len(matrix))
    for i, firstValue := range matrix {
        rows[i] = firstValue[0]
    }

    var rowIndex, left, right int;
    if len(matrix) > 1 {
        left, right = 0, len(rows)
        for left < right {
            var mid int = left + (right - left) / 2
            if rows[mid] == target {
                return true
            } else if rows[mid] > target {
                right = mid - 1
            } else {
                left = mid + 1
            }
        }
        rowIndex = left
    } else if len(matrix) == 1 {
        rowIndex = 0
    }

    if rowIndex < len(matrix) {
        left, right = 0, len(matrix[rowIndex])
        for left < right {
            var mid int = left + (right - left) / 2
            fmt.Printf("mid: %v\n", mid)
            fmt.Printf("matrix[rowIndex][mid]: %v\n", matrix[rowIndex][mid])
            if matrix[rowIndex][mid] == target {
                return true
            } else if matrix[rowIndex][mid] > target {
                right = mid
            } else {
                left = mid + 1
            }
        }
    }

	return false
}
