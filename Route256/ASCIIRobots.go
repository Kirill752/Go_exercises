package main

func whereToGo(field [][]byte, Ay int, Ax int, By int, Bx int, isA bool, isB bool) (bool, bool) {
	// Вниз вверх вправо влево
	// dirY := []int{1, -1, 0, 0}
	// dirX := []int{0, 0, 1, -1}
	dirYA := []int{-1, 0, 0, 1}
	dirXA := []int{0, -1, 1, 0}
	dirYB := []int{1, 0, -1, 0}
	dirXB := []int{0, 1, 0, -1}
	// Робот может пойти в 4 разных напрвления
	for i := 0; i < 4; i++ {
		// goA, goB := false, false
		if !isA {
			if field[Ay+dirYA[i]][Ax+dirXA[i]] == '.' {
				field[Ay+dirYA[i]][Ax+dirXA[i]] = 'a'
				// Если робот А достиг пункта назначения
				if Ay+dirYA[i] == 1 && Ax+dirXA[i] == 1 {
					isA = true
					return isA, isB
				}
				// Иначе робот А продолжает свой путь
				// Запускаем для его ного положения этот же алгоритм
				isA, isB = whereToGo(field, Ay+dirYA[i], Ax+dirXA[i], By, Bx, isA, isB)
				if isA && isB {
					return isA, isB
				}
			}
		}
		if !isB {
			if field[By+dirYB[i]][Bx+dirXB[i]] == '.' {
				field[By+dirYB[i]][Bx+dirXB[i]] = 'b'
				if By+dirYB[i] == (len(field)-2) && Bx+dirXB[i] == (len(field[0])-2) {
					isB = true
					return isA, isB
				}
				isA, isB = whereToGo(field, Ay, Ax, By+dirYB[i], Bx+dirXB[i], isA, isB)
				if isA && isB {
					return isA, isB
				}
			}
		}
	}
	return isA, isB
}

// i, j - индексы элемента, который упирается в никуда
func cleanField(field [][]byte, i int, j int) {
	// for i := 0; i < len(field); i++ {
	// 	fmt.Println(string(field[i]))
	// }
	// Вниз вверх вправо влево
	dirY := []int{1, -1, 0, 0}
	dirX := []int{0, 0, 1, -1}
	// Если элемент упирается в никуда
	// То есть, если три его соседа не он сам (не считаем конечные точки)
	// То превращаем его в точку
	neightborCount := 0
	if field[i][j] == 'a' {
		for k := 0; k < 4; k++ {
			if field[i+dirY[k]][j+dirX[k]] != 'a' && field[i+dirY[k]][j+dirX[k]] != 'A' {
				neightborCount++
			}
		}
		if neightborCount >= 3 {
			field[i][j] = '.'
			for k := 0; k < 4; k++ {
				if field[i+dirY[k]][j+dirX[k]] == 'a' {
					cleanField(field, i+dirY[k], j+dirX[k])
				}
			}
		}
	}
	neightborCount = 0
	if field[i][j] == 'b' {
		for k := 0; k < 4; k++ {
			if field[i+dirY[k]][j+dirX[k]] != 'b' && field[i+dirY[k]][j+dirX[k]] != 'B' {
				neightborCount++
			}
		}
		if neightborCount >= 3 {
			field[i][j] = '.'
			for k := 0; k < 4; k++ {
				if field[i+dirY[k]][j+dirX[k]] == 'b' {
					cleanField(field, i+dirY[k], j+dirX[k])
				}
			}
		}
	}

}
func ASCIIRobots(field [][]byte) [][]byte {
	coordsA := make([]int, 2)
	coordsB := make([]int, 2)
	copyField := make([][]byte, len(field))
	// Ищем место нахождение роботов A и B
	for i, str := range field {
		copyField[i] = make([]byte, len(field[i]))
		for j, ch := range str {
			if ch == 'A' {
				coordsA[0] = i
				coordsA[1] = j
			}
			if ch == 'B' {
				coordsB[0] = i
				coordsB[1] = j
			}
			copyField[i][j] = ch
		}
	}
	// Робот А должен попасть в клетку [0, 0]
	// Робот А должен попасть в клетку [n-1, m-1]
	isA, isB := whereToGo(field, coordsA[0], coordsA[1], coordsB[0], coordsB[1], false, false)
	if isA && isB {
		for i, str := range field {
			for j, ch := range str {
				if (i == 1 && j == 1) || (i == len(field)-2 && j == len(field[0])-2) {
					continue
				}
				if ch == 'a' || ch == 'b' {
					cleanField(field, i, j)
				}
			}
		}
		return field
	}
	return copyField
}
