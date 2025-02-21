package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var inOrder func(n *TreeNode)
	vals := []int{}
	res := true
	inOrder = func(n *TreeNode) {
		if !res {
			return
		}
		if n == nil {
			return
		}
		inOrder(n.left)
		if len(vals) == 0 {
			vals = append(vals, n.val)
		} else {
			if n.val > vals[len(vals)-1] {
				if len(vals) == 1 {
					vals = append(vals, n.val)
				} else {
					vals[0] = vals[1]
					vals[1] = n.val
				}
			} else {
				res = false
			}
		}
		inOrder(n.right)
	}
	inOrder(root)
	return res
}
