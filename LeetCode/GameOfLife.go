package main

func gameOfLife(board [][]int) {
	dj := []int{-1, 0, 1, 1, 1, 0, -1, -1}
	di := []int{-1, -1, -1, 0, 1, 1, 1, 0}
	// 2 - выживет в будущем
	// 3 - умрет в будущем
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 0 {
				aliveNeightbors := 0
				for m := 0; m < len(di); m++ {
					if i+di[m] > -1 && i+di[m] < len(board) && j+dj[m] > -1 && j+dj[m] < len(board[0]) && (board[i+di[m]][j+dj[m]] == 1 || board[i+di[m]][j+dj[m]] == 3) {
						aliveNeightbors++
					}
				}
				if aliveNeightbors == 3 {
					board[i][j] = 2
				}
			} else {
				if board[i][j] == 1 {
					aliveNeightbors := 0
					for m := 0; m < len(di); m++ {
						if i+di[m] > -1 && i+di[m] < len(board) && j+dj[m] > -1 && j+dj[m] < len(board[0]) && (board[i+di[m]][j+dj[m]] == 1 || board[i+di[m]][j+dj[m]] == 3) {
							aliveNeightbors++
						}
					}
					if aliveNeightbors < 2 || aliveNeightbors > 3 {
						board[i][j] = 3
					}
				}
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else {
				if board[i][j] == 3 {
					board[i][j] = 0
				}
			}
		}
	}
}
