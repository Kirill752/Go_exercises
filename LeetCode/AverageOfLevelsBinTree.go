package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{0}
	}
	average := []float64{}
	que := []*TreeNode{}
	que = append(que, root)
	Sum := float64(root.val)
	n := 1.0
	for len(que) > 0 {
		average = append(average, Sum/n)
		n, Sum = 0, 0
		L := len(que)
		for i := 0; i < L; i++ {
			if que[i].left != nil {
				que = append(que, que[i].left)
				Sum += float64(que[i].left.val)
				n++
			}
			if que[i].right != nil {
				que = append(que, que[i].right)
				Sum += float64(que[i].right.val)
				n++
			}
		}
		if len(que) > 0 {
			que = que[L:]
		}
	}
	return average
}
