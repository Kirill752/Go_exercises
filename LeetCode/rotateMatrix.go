package main

func rotateMatrix(matrix [][]int) {
	up, down := 0, len(matrix)-1
	for up < down {
		buf := matrix[up]
		matrix[up] = matrix[down]
		matrix[down] = buf
		up++
		down--
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j > i {
				buf := matrix[i][j]
				matrix[i][j] = matrix[j][i]
				matrix[j][i] = buf
			}
		}
	}

}
