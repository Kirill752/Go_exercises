package main

import "fmt"

func main() {
	pr := []int{1, 2, 5, 3, 4}
	in := []int{2, 5, 1, 3, 4}
	root := buildTree(pr, in)
	// root.left.left.val = 1
	fmt.Println(averageOfLevels(root))
}
