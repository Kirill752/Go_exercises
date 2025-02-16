package main

import "fmt"

func main() {
	t := InitBinTree([]int{3})
	t.root.Insesrt_in_Node(9)
	t.root.Insesrt_in_Node(20)
	t.root.right.Insesrt_in_Node(15)
	t.root.right.Insesrt_in_Node(7)

	t1 := InitBinTree([]int{3})
	t1.root.Insesrt_in_Node(9)
	t1.root.Insesrt_in_Node(20)
	t1.root.right.Insesrt_in_Node(16)
	t1.root.right.Insesrt_in_Node(7)
	fmt.Println(isSameTree(t.root, t1.root))
}
