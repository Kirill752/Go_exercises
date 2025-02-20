package main

import "fmt"

func main() {
	pr := []int{-1, 5, 4, 2, -4}
	in := []int{4, -4, 2, 5, -1}
	root := buildTree(pr, in)
	// root.left.left.val = 1
	fmt.Println(maxPathSum(root))
}
