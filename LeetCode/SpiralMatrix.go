package main

func spiralOrder(matrix [][]int) []int {
	xmin, ymin := 0, 0
	xmax, ymax := len(matrix[0])-1, len(matrix)-1
	spiral := []int{}
	for len(spiral) < len(matrix)*len(matrix[0]) {
		x, y := xmin, ymin
		for x <= xmax {
			spiral = append(spiral, matrix[y][x])
			x++
		}
		if len(spiral) >= len(matrix)*len(matrix[0]) {
			break
		}
		x--
		y++
		for y <= ymax {
			spiral = append(spiral, matrix[y][x])
			y++
		}
		if len(spiral) >= len(matrix)*len(matrix[0]) {
			break
		}
		y--
		x--

		xmax--
		ymin++
		for x >= xmin {
			spiral = append(spiral, matrix[y][x])
			x--
		}
		if len(spiral) >= len(matrix)*len(matrix[0]) {
			break
		}
		x++
		y--
		for y >= ymin {
			spiral = append(spiral, matrix[y][x])
			y--
		}
		if len(spiral) >= len(matrix)*len(matrix[0]) {
			break
		}
		xmin++
		ymax--
	}
	return spiral
}
