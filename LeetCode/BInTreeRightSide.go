package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) (canSee []int) {
	if root == nil {
		return []int{}
	}
	que := []*TreeNode{}
	que = append(que, root)
	for len(que) > 0 {
		canSee = append(canSee, que[len(que)-1].val)
		L := len(que)
		for i := 0; i < L; i++ {
			if que[i].left != nil {
				que = append(que, que[i].left)
			}
			if que[i].right != nil {
				que = append(que, que[i].right)
			}
		}
		if len(que) > 0 {
			que = que[L:]
		}
	}
	return
}
