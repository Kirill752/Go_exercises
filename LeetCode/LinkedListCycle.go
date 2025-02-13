package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == nil || fast.Next == nil {
			return false
		}
		if fast == slow || fast.Next == slow {
			return true
		}
	}
	return false
}
