package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	n := head
	l := 0
	for ; n.Next != nil; n = n.Next {
		l++
	}
	l++
	n.Next = head
	n = head
	for i := 1; i < l-(k%l); i++ {
		n = n.Next
	}
	head = n.Next
	n.Next = nil
	return head
}
