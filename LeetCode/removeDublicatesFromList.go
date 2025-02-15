package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	p := dummy
	c := head
	for c != nil {
		if c.Next != nil && c.Val == c.Next.Val {
			// Пропускаем значения
			for ; c.Next != nil && c.Val == c.Next.Val; c = c.Next {
			}
			p.Next = c.Next
		} else {
			p = p.Next
		}
		c = c.Next
	}
	return dummy.Next
}
