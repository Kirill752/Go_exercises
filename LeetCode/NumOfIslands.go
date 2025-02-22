package main

type cell struct {
	i, j int
}

// Не очень эффективно
func numIslandsBFS(grid [][]byte) int {
	numIslands := 0
	dx := []int{0, -1, 0, 1}
	dy := []int{-1, 0, 1, 0}
	seen := make(map[cell]bool)
	que := []cell{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			c := cell{i, j}
			if grid[i][j] == '1' && !seen[c] {
				// Добавляем в очередь
				que = append(que, c)
				for len(que) > 0 {
					// Извлекаем первый элемент из очереди
					cur := que[0]
					que = que[1:]
					seen[cur] = true
					// Добавляем в очередь всех соседей == 1
					for k, di := range dx {
						dj := dy[k]
						if !seen[cell{cur.i + di, cur.j + dj}] && (cur.i+di >= 0 && cur.i+di < len(grid)) && (cur.j+dj >= 0 && cur.j+dj < len(grid[0])) {
							if grid[cur.i+di][cur.j+dj] == '1' {
								que = append(que, cell{cur.i + di, cur.j + dj})
							}
						}
					}
				}
				numIslands++
			}
		}
	}
	return numIslands
}

// Лучше использовать это
func numIslandsDFS(grid [][]byte) int {
	numIslands := 0
	dx := []int{0, -1, 0, 1}
	dy := []int{-1, 0, 1, 0}
	seen := make(map[cell]bool)
	var inOrder func(g [][]byte, i, j int)
	inOrder = func(g [][]byte, i, j int) {
		cur := cell{i, j}
		seen[cur] = true
		for k, di := range dx {
			dj := dy[k]
			if !seen[cell{cur.i + di, cur.j + dj}] && (cur.i+di >= 0 && cur.i+di < len(grid)) && (cur.j+dj >= 0 && cur.j+dj < len(grid[0])) {
				if grid[cur.i+di][cur.j+dj] == '1' {
					inOrder(g, cur.i+di, cur.j+dj)
				}
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !seen[cell{i, j}] {
				inOrder(grid, i, j)
				numIslands++
			}
		}
	}
	return numIslands
}
