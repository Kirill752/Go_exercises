package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	start *TreeNode
}

func ConstructorB(root *TreeNode) BSTIterator {
	var iter BSTIterator
	iter.start = new(TreeNode)
	n := root
	for ; n.left != nil; n = n.left {
	}
	iter.start.val = n.val - 1
	N := iter.start
	var inOrder func(*TreeNode)
	inOrder = func(n *TreeNode) {
		if n == nil {
			return
		}
		inOrder(n.left)
		N.right = &TreeNode{val: n.val}
		N = N.right
		inOrder(n.right)
	}
	inOrder(root)
	return iter
}

func (this *BSTIterator) Next() int {
	buf := this.start.right.val
	this.start = this.start.right
	return buf
}

func (this *BSTIterator) HasNext() bool {
	return this.start.right != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
