package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSumforNode(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var maxSum, sum int
	var preorder func(n *TreeNode, sum int)
	preorder = func(n *TreeNode, sum int) {
		if n == nil {
			return
		}
		sum += n.val
		maxSum = max(maxSum, sum)
		if n.left == nil && n.right == nil {
			sum -= n.val
			return
		}
		preorder(n.left, sum)
		preorder(n.right, sum)
	}
	maxSum = -100000000
	sum = 0
	preorder(root.left, sum)
	maxL := maxSum
	maxSum = -1000000000
	sum = 0
	preorder(root.right, sum)
	maxR := maxSum
	return max(root.val, maxL, maxR, maxL+root.val, maxR+root.val, maxL+maxR+root.val)
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res = -100000000
	var preorder func(n *TreeNode)
	preorder = func(n *TreeNode) {
		if n == nil {
			return
		}
		res = max(res, maxPathSumforNode(n))
		preorder(n.left)
		preorder(n.right)
	}
	preorder(root)
	return res
}
