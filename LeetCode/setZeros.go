package main

func setZeroes(matrix [][]int) {
	zerolines := make([]int, len(matrix))
	zerocoloums := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				zerolines[i] = 1
				zerocoloums[j] = 1
			}
		}
	}
	for i := 0; i < len(zerolines); i++ {
		if zerolines[i] == 1 {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 0; j < len(zerocoloums); j++ {
		if zerocoloums[j] == 1 {
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
}
