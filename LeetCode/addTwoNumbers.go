package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	buf := 0
	r := new(ListNode)
	start := r
	n1, n2 := l1, l2
	for ; n1 != nil || n2 != nil; n1, n2 = n1.Next, n2.Next {
		if n1 == nil {
			n1 = new(ListNode)
		}
		if n2 == nil {
			n2 = new(ListNode)
		}
		res := (n1.Val + n2.Val + buf)
		buf = res / 10
		r.Val = res % 10
		if n1.Next == nil && n2.Next == nil {
			break
		}
		r.Next = new(ListNode)
		r = r.Next
	}
	if buf != 0 {
		r.Next = new(ListNode)
		r.Next.Val = buf
	}
	return start
}
