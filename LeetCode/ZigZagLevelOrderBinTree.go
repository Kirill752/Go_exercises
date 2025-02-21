package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
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
			if len(levelOrder)%2 == 0 {
				for i := 0; i < len(que); i++ {
					newlevel[i] = que[i].val
				}
			} else {
				for i := len(que) - 1; i > -1; i-- {
					newlevel[(len(newlevel)-1)-i] = que[i].val
				}
			}
			levelOrder = append(levelOrder, newlevel)
		}
	}
	return levelOrder
}
