package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	inf = 1 << 30
)

var (
	n, m, d int
	grid    [][]byte
	dist    [][]int
	dx      = []int{-1, 1, 0, 0}
	dy      = []int{0, 0, -1, 1}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d %d\n", &n, &m, &d)
	grid = make([][]byte, n)
	for i := range n {
		grid[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			grid[i][j], _ = reader.ReadByte()
		}
		reader.ReadByte() // Пропускаем символ новой строки
	}

	// Инициализация матрицы расстояний
	dist = make([][]int, n)
	for i := range n {
		dist[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if grid[i][j] == 'x' {
				dist[i][j] = 0
			} else {
				dist[i][j] = inf
			}
		}
	}

	// BFS для заполнения расстояний
	queue := make([][2]int, 0)
	for i := range n {
		for j := range m {
			if grid[i][j] == 'x' {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		x, y := cell[0], cell[1]
		for k := range 4 {
			nx, ny := x+dx[k], y+dy[k]
			if nx >= 0 && nx < n && ny >= 0 && ny < m && dist[nx][ny] == inf {
				dist[nx][ny] = dist[x][y] + 1
				queue = append(queue, [2]int{nx, ny})
			}
		}
	}

	// Бинарный поиск по размеру квадрата
	left, right := 0, min(n, m)
	maxSize := 0
	for left <= right {
		mid := (left + right) / 2
		if canPlaceSquare(mid) {
			maxSize = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Println(maxSize)
}

func canPlaceSquare(k int) bool {
	if k == 0 {
		return true
	}
	for i := 0; i <= n-k; i++ {
		for j := 0; j <= m-k; j++ {
			if dist[i][j] >= d && dist[i+k-1][j] >= d && dist[i][j+k-1] >= d && dist[i+k-1][j+k-1] >= d {
				return true
			}
		}
	}
	return false
}
