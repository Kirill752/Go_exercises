package main

// func findWord(board [][]byte, word string, i, j int) {
// 	di := []int{0, -1, 0, 1}
// 	dj := []int{-1, 0, 1, 0}
// 	seen := map[cell]bool{}
// 	curr := cell{i, j}
// 	seen[curr] = true
// 	res := false
// 	var search func(c cell) bool
// 	search = func(c cell) bool {
// 		for i := 0; i < len(word); i++ {
// 			ch := word[i]
// 			if board[curr.i][curr.j] != ch {
// 				return false
// 			}
// 			for m := 0; m < len(di); m++ {
// 				if i+di[m] > -1 && i+di[m] < len(board) && j+dj[m] > -1 && j+dj[m] < len(board[0]) {
// 					c := cell{di[m], dj[m]}
// 					if !seen[c] {
// 						seen[c] = true
// 						curr = c
// 						search(c)
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// func findWords(board [][]byte, words []string) []string {
// 	t := ConstructorTrie()
// 	// Делаем префиксное дерево из всех возможных комбинация слов
// 	for i := 0; i < len(board); i++ {
// 		for j := 0; j < len(board[0]); j++ {
// 			t.Insert(string(board[i][j]))
// 		}
// 	}
// }
