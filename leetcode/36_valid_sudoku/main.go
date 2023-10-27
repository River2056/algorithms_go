package validsudoku

import "fmt"

func validate(arr []byte) bool {
	m := make(map[byte]int)
    for _, v := range arr {
        fmt.Printf("%v ", string(v))
    }
    fmt.Println()
	for _, b := range arr {
        if b == '.' {
            continue
        }
		if _, ok := m[b]; ok {
			return false
		} else {
			m[b] = 1
		}
	}

	return true
}

func IsValidSudoku(board [][]byte) bool {
	var isValid bool
	for _, row := range board {
		isValid = validate(row)
		if !isValid {
			return false
		}
	}
    fmt.Println()

	for row, _ := range board[0] {
		column := make([]byte, 0)
		for col, _ := range board {
			column = append(column, board[col][row])
		}
        isValid = validate(column)
        if !isValid {
            return false
        }
	}
    fmt.Println()

	grid := make(map[string][]byte)

	for r, row := range board {
		for c := range row {
			key := fmt.Sprintf("%d,%d", r / 3, c / 3)
			if _, ok := grid[key]; !ok {
				grid[key] = make([]byte, 0)
			}
			grid[key] = append(grid[key], board[r][c])
		}
	}

	for _, g := range grid {
		isValid = validate(g)
		if !isValid {
			return false
		}
	}

	return true
}
