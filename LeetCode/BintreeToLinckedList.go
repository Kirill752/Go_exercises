package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		return
	}
	flatten(root.left)
	flatten(root.right)
	if root.left != nil {
		buf := root.right
		root.right = root.left
		root.left = nil
		r := root
		for ; r.right != nil; r = r.right {
		}
		r.right = buf
	}
}
