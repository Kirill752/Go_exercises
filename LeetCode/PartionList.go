package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertNode(n *ListNode, i *ListNode) {
	i.Next = n.Next
	n.Next = i
}
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := new(ListNode)
	dummy.Next = head
	n := dummy
	toIns := dummy
	for ; n.Next != nil && n.Next.Val < x; n = n.Next {
	}
	toIns = n
	for n != nil && n.Next != nil {
		if n.Next.Val < x {
			goTo := n.Next.Next
			insertNode(toIns, n.Next)
			toIns = toIns.Next
			n.Next = goTo
		} else {
			n = n.Next
		}
	}
	return dummy.Next
}
