package main

func solve(board [][]byte) {
	dx := []int{0, -1, 0, 1}
	dy := []int{-1, 0, 1, 0}
	seen := make(map[cell]bool)
	toChange := make(map[cell]bool)
	var inOrder func(g [][]byte, i, j int)
	marker := true
	inOrder = func(g [][]byte, i, j int) {
		cur := cell{i, j}
		seen[cur] = true
		toChange[cur] = true
		for k, di := range dx {
			dj := dy[k]
			if !seen[cell{cur.i + di, cur.j + dj}] {
				if (cur.i+di >= 0 && cur.i+di < len(board)) && (cur.j+dj >= 0 && cur.j+dj < len(board[0])) {
					if board[cur.i+di][cur.j+dj] == 'O' {
						inOrder(g, cur.i+di, cur.j+dj)
					}
				} else {
					marker = false
				}
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' && !seen[cell{i, j}] {
				inOrder(board, i, j)
				if marker {
					for c := range toChange {
						board[c.i][c.j] = 'X'
					}
				} else {
					marker = true
				}
				clear(toChange)
			}
		}
	}
}
