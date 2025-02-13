package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var res *ListNode
	if list1.Val < list2.Val {
		res = list1
		list1 = list1.Next
	} else {
		res = list2
		list2 = list2.Next
	}
	r := res
	for n1, n2 := list1, list2; n1 != nil || n2 != nil; {
		if n1 == nil {
			res.Next = n2
			n2 = n2.Next
			res = res.Next
			continue
		}
		if n2 == nil {
			res.Next = n1
			n1 = n1.Next
			res = res.Next
			continue
		}
		if n1.Val < n2.Val {
			res.Next = n1
			n1 = n1.Next
		} else {
			res.Next = n2
			n2 = n2.Next
		}
		res = res.Next
	}
	return r
}
