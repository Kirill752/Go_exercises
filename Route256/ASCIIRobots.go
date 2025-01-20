package main

func whereToGo(field [][]byte, Ay int, Ax int, By int, Bx int, isA bool, isB bool) (bool, bool) {
	// Вниз вверх вправо влево
	// dirY := []int{1, -1, 0, 0}
	// dirX := []int{0, 0, 1, -1}
	dirYA := []int{-1, 0}
	dirXA := []int{0, -1}
	dirYB := []int{1, 0}
	dirXB := []int{0, 1}
	if Ay == 1 && Ax == 1 {
		isA = true
	}
	if By == (len(field)-2) && Bx == (len(field[0])-2) {
		isB = true
	}
	// Робот может пойти в 4 разных напрвления
	for i := 0; i < 2; i++ {
		// goA, goB := false, false
		if !isA {
			if field[Ay+dirYA[i]][Ax+dirXA[i]] == '.' {
				field[Ay+dirYA[i]][Ax+dirXA[i]] = 'a'
				// Если робот А достиг пункта назначения
				if Ay+dirYA[i] == 1 && Ax+dirXA[i] == 1 {
					isA = true
					if isA && isB {
						return isA, isB
					}
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
					if isA && isB {
						return isA, isB
					}
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
	// Определяем, какого робота толкать влево, а какого вправо
	switch coordsA[0]+coordsA[1] < coordsB[0]+coordsB[1] {
	// толкаем робота A влево-вверх, а робота B вправо-вниз
	case true:
		isA, isB := whereToGo(field, coordsA[0], coordsA[1], coordsB[0], coordsB[1], false, false)
		if isA && isB {
			return field
		}
	case false:
		field[coordsA[0]][coordsA[1]] = 'B'
		field[coordsB[0]][coordsB[1]] = 'A'
		isA, isB := whereToGo(field, coordsB[0], coordsB[1], coordsA[0], coordsA[1], false, false)
		if isA && isB {
			field[coordsA[0]][coordsA[1]] = 'A'
			field[coordsB[0]][coordsB[1]] = 'B'
			for i, str := range field {
				for j, ch := range str {
					if ch == 'a' {
						field[i][j] = 'b'
					} else {
						if ch == 'b' {
							field[i][j] = 'a'
						}
					}
				}
			}
			return field
		}
	}
	return copyField
}
