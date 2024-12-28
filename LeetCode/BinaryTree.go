package main

import (
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

type Tree struct {
	root *TreeNode
}

/*Вставка элемента в дерево*/
func (t *Tree) Insesrt(val int) {
	// Проверка существования дерева
	if t.root == nil {
		t.root = &TreeNode{val: val}
		return
	}
	// Начинаем обход с корня
	current := t.root
	for {
		// Если такое значение уже существует, то возвращаем ошибку
		if current.val == val {
			fmt.Println("This node value has already exists!")
			return
		}
		// Если вставляемое значение меньше текущего, то переходим влево
		if current.val > val {
			if current.left == nil {
				current.left = &TreeNode{val: val}
				return
			}
			current = current.left
		}
		// Если вставляемое значение больше текущего, то переходим вправо
		if current.val < val {
			if current.right == nil {
				current.right = &TreeNode{val: val}
				return
			}
			current = current.right
		}
	}
}

func (t *TreeNode) Insesrt_in_Node(val int) {
	if t.left == nil {
		t.left = &TreeNode{val: val}
		return
	}
	if t.right == nil {
		t.right = &TreeNode{val: val}
		return
	}
	if (t.left.left != nil) && (t.left.right != nil) {
		t.right.Insesrt_in_Node(val)
	} else {
		t.left.Insesrt_in_Node(val)
	}
}

// Печать дерева в консоль
func PrintTree(Node *TreeNode, space int) {
	if Node != nil {
		PrintTree(Node.right, space+5)
		for i := 0; i < space; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("%d\n", Node.val)
		PrintTree(Node.left, space+5)
	}
}
func PrintTree_2(Node *TreeNode, rpref string, cpref string, lpref string) {
	if Node == nil {
		return
	}
	if Node.right != nil {
		PrintTree_2(Node.right, rpref+" ", rpref+"/-", rpref+"| ")
	}
	fmt.Printf("%s%d\n", cpref, Node.val)
	if Node.left != nil {
		PrintTree_2(Node.left, lpref+"| ", lpref+"\\-", lpref+" ")
	}
}

func largestValues(root []int) []int {
	node := &TreeNode{val: root[0]}
	for i := 1; i < len(root); i++ {
		node.Insesrt_in_Node(root[i])
	}
	PrintTree_2(node, "", "", "")
	return nil
}
