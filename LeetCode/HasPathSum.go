package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.val
	if root.left == nil && root.right == nil {
		return targetSum == 0
	}
	if root.left == nil {
		return hasPathSum(root.right, targetSum)
	}
	if root.right == nil {
		return hasPathSum(root.left, targetSum)
	}
	return hasPathSum(root.left, targetSum) || hasPathSum(root.right, targetSum)
}
