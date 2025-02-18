package main

import "fmt"

func main() {
	pr := []int{1, 2, 3, 4, 5, 6}
	in := []int{3, 2, 4, 1, 5, 6}
	root := buildTree(pr, in)
	fmt.Println(hasPathSum(root, 7))
}
