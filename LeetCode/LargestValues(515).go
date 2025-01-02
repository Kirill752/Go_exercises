package main

import (
	"fmt"
	"math"
)

type pair struct {
	ptr   *TreeNode
	level int
}

func remove(slice []pair, i int) []pair {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
func largestValues_1(root []int) []int {
	node := &TreeNode{val: root[0]}
	for i := 1; i < len(root); i++ {
		node.Insesrt_in_Node(root[i])
	}
	max_level := int(math.Round(math.Log2(float64(len(root) + len(root)%2))))
	result := make([]int, max_level)
	deq := []pair{}
	deq = append(deq, pair{ptr: node, level: 0})
	for len(deq) > 0 {
		result[deq[0].level] = max(result[deq[0].level], deq[0].ptr.val)
		if deq[0].ptr.left != nil {
			deq = append(deq, pair{ptr: deq[0].ptr.left, level: (deq[0].level + 1)})
		}
		if deq[0].ptr.right != nil {
			deq = append(deq, pair{ptr: deq[0].ptr.right, level: (deq[0].level + 1)})
		}
		deq = remove(deq, 0)
	}
	return result
}
func largestValues(root *TreeNode) []int {
	res_map := make(map[int]int)
	deq := []pair{}
	deq = append(deq, pair{ptr: root, level: 0})
	for len(deq) > 0 {
		if _, ok := res_map[deq[0].level]; ok {
			res_map[deq[0].level] = max(res_map[deq[0].level], deq[0].ptr.val)
		} else {
			res_map[deq[0].level] = deq[0].ptr.val
		}
		if deq[0].ptr.left != nil {
			deq = append(deq, pair{ptr: deq[0].ptr.left, level: (deq[0].level + 1)})
		}
		if deq[0].ptr.right != nil {
			deq = append(deq, pair{ptr: deq[0].ptr.right, level: (deq[0].level + 1)})
		}
		deq = remove(deq, 0)
	}
	result := make([]int, len(res_map))
	fmt.Println(res_map)
	for i := 0; i < len(res_map); i++ {
		result[i] = res_map[i]
	}
	return result
}
