package main

// type light struct {
// 	x int
// 	y int
// 	// D, U, L, R
// 	dir byte
// }

func givelight(room [][]int, x, y int, dir byte, totalLight int) int {
	if (x > -1 && x < len(room)) && (y > -1 && y < len(room[0])) {
		if room[x][y] == 1 {
			return totalLight
		}
		room[x][y] = 1
		switch dir {
		case 'D':
			totalLight += givelight(room, x+1, y-1, dir, totalLight)
			totalLight += givelight(room, x+1, y, dir, totalLight)
			totalLight += givelight(room, x+1, y+1, dir, totalLight)
		case 'U':
			totalLight += givelight(room, x-1, y-1, dir, totalLight)
			totalLight += givelight(room, x-1, y, dir, totalLight)
			totalLight += givelight(room, x-1, y+1, dir, totalLight)
		case 'L':
			totalLight += givelight(room, x-1, y-1, dir, totalLight)
			totalLight += givelight(room, x, y-1, dir, totalLight)
			totalLight += givelight(room, x+1, y-1, dir, totalLight)
		case 'R':
			totalLight += givelight(room, x-1, y+1, dir, totalLight)
			totalLight += givelight(room, x, y+1, dir, totalLight)
			totalLight += givelight(room, x+1, y+1, dir, totalLight)
		}
	}
	return totalLight
}
func bestDir(room [][]int, x, y int) byte {
	// case D
	freeD := 0
	if x+1 > -1 && x+1 < len(room) {
		for j := y - 1; j <= y+1; j++ {
			if j < 0 || j >= len(room[0]) {
				continue
			}
			if room[x+1][j] == 0 {
				freeD++
			}
		}
	}
	// case U
	freeU := 0
	if x-1 > -1 && x-1 < len(room) {
		for j := y - 1; j <= y+1; j++ {
			if j < 0 || j >= len(room[0]) {
				continue
			}
			if room[x-1][j] == 0 {
				freeU++
			}
		}
	}
	free := max(freeD, freeU)
	// case L
	freeL := 0
	if y-1 > -1 && y-1 < len(room[0]) {
		for i := x - 1; i <= x+1; i++ {
			if i < 0 || i >= len(room) {
				continue
			}
			if room[i][y-1] == 0 {
				freeL++
			}
		}
	}
	free = max(free, freeL)
	// case R
	freeR := 0
	if y+1 > -1 && y+1 < len(room[0]) {
		for i := x - 1; i <= x+1; i++ {
			if i < 0 || i >= len(room) {
				continue
			}
			if room[i][y+1] == 0 {
				freeR++
			}
		}
	}
	free = max(free, freeR)
	switch free {
	case freeD:
		return 'D'
	case freeU:
		return 'U'
	case freeL:
		return 'L'
	case freeR:
		return 'R'
	}
	return 'D'
}
func darkRoom(room [][]int) ([]int, []int, []byte) {
	l := []int{}
	k := []int{}
	dirs := []byte{}
	totalLight := 0
	for x := 0; x < len(room) && totalLight <= len(room)*len(room[0]); x++ {
		for y := 0; y < len(room[0]) && totalLight <= len(room)*len(room[0]); y++ {
			if room[x][y] == 0 {
				dir := bestDir(room, x, y)
				l = append(l, x)
				k = append(k, y)
				dirs = append(dirs, dir)
				totalLight = givelight(room, x, y, dir, totalLight)
			}
		}
	}
	return l, k, dirs
}
