package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	levelOrder := [][]int{}
	que := []*TreeNode{}
	que = append(que, root)
	levelOrder = append(levelOrder, []int{root.val})
	for len(que) > 0 {
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
		if len(que) > 0 {
			newlevel := make([]int, len(que))
			for i, v := range que {
				newlevel[i] = v.val
			}
			levelOrder = append(levelOrder, newlevel)
		}
	}
	return levelOrder
}
