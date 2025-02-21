package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	minDiff := 100000
	p, c := -1, -1
	var inOrder func(n *TreeNode)
	inOrder = func(n *TreeNode) {
		if n == nil {
			return
		}
		inOrder(n.left)
		switch c {
		case -1:
			c = n.val
		default:
			p = c
			c = n.val
		}
		if p != -1 {
			minDiff = min(minDiff, c-p)
		}
		inOrder(n.right)
	}
	inOrder(root)
	return minDiff
}
