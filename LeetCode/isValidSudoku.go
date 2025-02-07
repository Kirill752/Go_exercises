package main

func isValidSudoku(board [][]byte) bool {
	digits := map[byte]int{'1': 0, '2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8}
	for _, line := range board {
		nums := make([]byte, 9)
		for _, digit := range line {
			if digit != '.' {
				nums[digits[digit]]++
				if nums[digits[digit]] > 1 {
					return false
				}
			}
		}
	}
	for coloumn := range board[0] {
		nums := make([]byte, 9)
		for j := 0; j < 9; j++ {
			if board[j][coloumn] != '.' {
				nums[digits[board[j][coloumn]]]++
				if nums[digits[board[j][coloumn]]] > 1 {
					return false
				}
			}
		}
	}
	di := []int{0, 1, 1, 1, 0, -1, -1, -1, 0}
	dj := []int{0, -1, 0, 1, 1, 1, 0, -1, -1}
	for i := 1; i < 8; i += 3 {
		for j := 1; j < 8; j += 3 {
			nums := make([]byte, 9)
			for k := 0; k < len(di); k++ {
				if board[i+di[k]][j+dj[k]] != '.' {
					nums[digits[board[i+di[k]][j+dj[k]]]]++
					if nums[digits[board[i+di[k]][j+dj[k]]]] > 1 {
						return false
					}
				}
			}
		}
	}
	return true
}
