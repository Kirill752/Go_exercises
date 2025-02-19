package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func StoI(S []int) (I int) {
	degree := 1
	for i := len(S) - 1; i > -1; i-- {
		I += S[i] * degree
		degree *= 10
	}
	return
}

func sumNumbers(root *TreeNode) int {
	s := 0
	number := []int{}
	var preorder func(n *TreeNode)
	preorder = func(n *TreeNode) {
		if n == nil {
			return
		}
		number = append(number, n.val)
		preorder(n.left)
		preorder(n.right)
		if n.left == nil && n.right == nil {
			s += StoI(number)
		}
		number = number[:len(number)-1]
	}
	preorder(root)
	return s
}
