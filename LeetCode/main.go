package main

import "fmt"

func main() {
	pr := []int{4, 9, 5, 1, 0}
	in := []int{5, 9, 1, 4, 0}
	root := buildTree(pr, in)
	fmt.Println(sumNumbers(root))
}
