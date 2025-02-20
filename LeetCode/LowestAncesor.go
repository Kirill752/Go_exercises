package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func Ansector(root, child *TreeNode) (ans *TreeNode) {
	ans = root
	flag := false
	var inOrder func(*TreeNode)
	inOrder = func(n *TreeNode) {
		if n == nil {
			return
		}
		if n.left == child || n.right == child {
			ans = n
			flag = true
			return
		}
		if flag {
			return
		}
		inOrder(n.left)
		inOrder(n.right)
	}
	inOrder(root)
	return
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pSeen := make(map[*TreeNode]bool)
	qSeen := make(map[*TreeNode]bool)
	for q != p && !pSeen[q] && !qSeen[p] {
		pSeen[p] = true
		qSeen[q] = true
		p = Ansector(root, p)
		q = Ansector(root, q)
	}
	if pSeen[q] {
		return q
	}
	if qSeen[p] {
		return p
	}
	return p
}
