package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.left, root.right = root.right, root.left
	invertTree(root.left)
	invertTree(root.right)
	return root
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	invertTree(root.left)
	return isSameTree(root.left, root.right)
}
