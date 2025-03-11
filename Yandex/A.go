package main

func notSleep(building [][]rune, n, m, x, y int) int {
	di := make([]int, 0, x*y)
	dj := make([]int, 0, x*y)
	for i := range y {
		for j := range x {
			di = append(di, i)
			dj = append(dj, j)
		}
	}
	res := 0
	for i := 0; i < n*x; i += x {
		for j := 0; j < m*y; j += y {
			cnt := 0
			for l := range di {
				if i+di[l] > -1 && i+di[l] < n*x &&
					j+dj[l] > -1 && j+dj[l] < m*y &&
					building[i+di[l]][j+dj[l]] == 'X' {
					cnt++
				}
			}
			if cnt >= (x*y)/2+(x*y)%2 {
				res++
			}
		}
	}
	return res
}

// func main() {
// 	n, m, x, y := 2, 1, 1, 1
// 	building := [][]rune{
// 		{'0'}, {'X'},
// 	}
// 	fmt.Println(notSleep(building, n, m, x, y))
// }

// Tests:
//
// {'X', '0', '0', '0'},
// {'0', '0', '0', 'X'},
// {'X', '0', '0', '0'},
// {'X', 'X', '0', '0'}
//
// {'0'},
// {'0'},
// {'X'},
// {'X'},
// {'0'},
// {'X'},
