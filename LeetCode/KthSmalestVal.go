package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var smalestVal int
	var inOrder func(n *TreeNode)
	inOrder = func(n *TreeNode) {
		if n == nil {
			return
		}
		inOrder(n.left)
		if k == 0 {
			return
		}
		smalestVal = n.val
		k--
		inOrder(n.right)
	}
	inOrder(root)
	return smalestVal
}
