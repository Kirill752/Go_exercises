package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t string
	t, _ = in.ReadString('\n')
	tINT, _ := strconv.Atoi(t[:len(t)-1])
	for i := 0; i < tINT; i++ {
		n, _ := in.ReadString('\n')
		nINT, _ := strconv.Atoi(n[:len(n)-1])
		// считываем массив arrival
		arraivalTable := make(map[int][]int)
		var arraival_i int
		for j := 0; j < nINT; j++ {
			fmt.Fscan(in, &arraival_i)
			arraivalTable[arraival_i] = append(arraivalTable[arraival_i], j)
		}
		in.ReadString('\n')
		m, _ := in.ReadString('\n')
		mINT, _ := strconv.Atoi(m[:len(m)-1])
		benzTrucks := make([]truck, mINT)
		for k := 0; k < mINT; k++ {
			benzTrucks[k].id = k + 1
			fmt.Fscan(in, &benzTrucks[k].start)
			fmt.Fscan(in, &benzTrucks[k].end)
			fmt.Fscan(in, &benzTrucks[k].capacity)
		}
		in.ReadString('\n')
		ans := planingOrder(nINT, arraivalTable, mINT, benzTrucks)
		for _, v := range ans {
			fmt.Fprintf(out, "%d ", v)
		}
		fmt.Fprintf(out, "\n")
		// fmt.Println(arraivalTable)
		// fmt.Println(benzTrucks)
	}
	// n := 1
	// arraival := []int{1000000000}
	// arraivalTable := make(map[int][]int)
	// for i := 0; i < len(arraival); i++ {
	// 	arraivalTable[arraival[i]] = append(arraivalTable[arraival[i]], i)
	// }
	// fmt.Println(arraivalTable)
	// m := 4
	// benzTrucks := make([]truck, m)
	// benzTrucks[0] = truck{1, 1, 1000000000, 1}
	// benzTrucks[1] = truck{2, 1, 46, 2}
	// benzTrucks[2] = truck{3, 41, 83, 1}
	// benzTrucks[3] = truck{4, 4, 75, 2}
	// fmt.Println(benzTrucks)
	// answer := planingOrder(n, arraivalTable, m, benzTrucks)
	// fmt.Println(answer)
	// slices.SortStableFunc(benzTrucks, func(a truck, b truck) int {
	// 	if a.start > b.start {
	// 		return 1
	// 	}
	// 	if a.start == b.start {
	// 		return 0
	// 	}
	// 	return -1
	// })
	// fmt.Println(benzTrucks)
}
