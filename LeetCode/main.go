package main

import "fmt"

func main() {
	pr := []int{7, 3, 15, 9, 20}
	in := []int{3, 7, 9, 15, 20}
	root := buildTree(pr, in)
	// root.left.left.val = 1
	fmt.Println(lowestCommonAncestor(root, root.right.left, root.left).val)
}
