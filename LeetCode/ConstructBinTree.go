package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.val = preorder[0]
	i := 0
	for _, v := range inorder {
		if v == root.val {
			break
		}
		i++
	}
	lenLeft := i
	// lenRight := len(inorder) - i + 1
	// В preorder левое дерево - это все от [1, lenLeft],
	// а правое дерво - это все [lenleft+1, len(preorder)-1]

	// В inorder левое дерево - это все от [0, lenLeft-1],
	// а правое дерво - это все [lenleft+1, len(inorder)-1]

	// Таким образом
	root.left = buildTree(preorder[1:lenLeft+1], inorder[:lenLeft])
	root.right = buildTree(preorder[lenLeft+1:], inorder[lenLeft+1:])

	// fmt.Printf("preorder: %v\n", preorder)
	// fmt.Printf("inorder: %v\n", inorder)
	// fmt.Println(inorder[i])
	return root
}
