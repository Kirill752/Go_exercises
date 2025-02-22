package main

import "fmt"

func main() {
	// pr := []int{1, 2, 5, 3, 4}
	// in := []int{2, 5, 1, 3, 4}
	// root := buildTree(pr, in)
	// root.left.left.val = 1
	board := [][]byte{{'O', 'X', 'O'}, {'X', 'O', 'X'}, {'O', 'X', 'O'}}
	for _, l := range board {
		fmt.Println(string(l))
	}
	fmt.Println()
	solve(board)
	for _, l := range board {
		fmt.Println(string(l))
	}
}
