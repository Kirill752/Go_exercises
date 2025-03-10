package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	res := lists[0]
	for i := 1; i < len(lists); i++ {
		res = mergeTwoLists(res, lists[i])
	}
	return res
}
