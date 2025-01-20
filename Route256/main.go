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
		field := make([][]byte, n+2)
		for i := 0; i < n+2; i++ {
			field[i] = make([]byte, m+2)
		}
		// Ограничиваем поле сверху и снизу
		for j := 0; j < m+2; j++ {
			field[0][j] = '#'
			field[n+1][j] = '#'
		}
		// Ограничиваем поле слева и справа
		for i := 0; i < n+2; i++ {
			field[i][0] = '#'
			field[i][m+1] = '#'
		}
		for i := 1; i < n+1; i++ {
			str, _ := in.ReadBytes('\n')
			for j := 1; j < m+1; j++ {
				field[i][j] = str[j-1]
			}
		}
		res := ASCIIRobots(field)
		fmt.Fprintln(out, "Answer")
		for i := 1; i < n+1; i++ {
			fmt.Fprintln(out, string(res[i][1:len(res[i])-1]))
		}
	}
}
