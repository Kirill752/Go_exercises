package main

// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func UpdateDeth(n *TreeNode, d int) int {
	if n.left == nil && n.right == nil {
		return d
	}
	var dl, dr int = d, d
	if n.left != nil {
		dl = UpdateDeth(n.left, dl+1)
	}
	if n.right != nil {
		dr = UpdateDeth(n.right, dr+1)
	}
	return max(dl, dr)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// maxD := UpdateDeth(root, 1)
	return 1 + max(maxDepth(root.left), maxDepth(root.right))
}
