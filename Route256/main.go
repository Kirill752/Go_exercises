package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	fmt.Fscanf(in, "%d", &t)
	_, _ = in.ReadByte()
	for ; t > 0; t-- {
		var n int // число строк
		var m int // число столбцов
		fmt.Fscanf(in, "%d %d", &n, &m)
		_, _ = in.ReadByte()
		room := make([][]int, n)
		for i := 0; i < n; i++ {
			room[i] = make([]int, m)
		}
		lenLight := 0
		totalLight := 0
		for x := 0; x < n && totalLight <= n*m; x++ {
			for y := 0; y < m && totalLight <= n*m; y++ {
				if room[x][y] == 0 {
					dir := bestDir(room, x, y)
					lenLight++
					totalLight = givelight(room, x, y, dir, totalLight)
				}
			}
		}
		for x := 0; x < n; x++ {
			for y := 0; y < m; y++ {
				room[x][y] = 0
			}
		}
		fmt.Fprintln(out, lenLight)
		totalLight = 0
		for x := 0; x < n && totalLight <= n*m; x++ {
			for y := 0; y < m && totalLight <= n*m; y++ {
				if room[x][y] == 0 {
					dir := bestDir(room, x, y)
					fmt.Fprintln(out, x+1, y+1, string(dir))
					totalLight = givelight(room, x, y, dir, totalLight)
				}
			}
		}
	}
}
